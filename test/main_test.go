package test

import (
	"backend-core/common"
	"backend-core/domain"
	"context"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/go-resty/resty/v2"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	tc "github.com/testcontainers/testcontainers-go"
	"log"
	"net/http"
	"testing"
)

func TestUser(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var _ = Describe("Main", func() {
	var ctx = context.Background()
	var network = tc.NetworkRequest{
		Name:   "integration-test-network",
		Driver: "bridge",
	}

	provider, err := tc.NewDockerProvider()
	if err != nil {
		log.Fatal(err)
	}

	if _, err := provider.GetNetwork(ctx, network); err != nil {
		if _, err := provider.CreateNetwork(ctx, network); err != nil {
			log.Fatal(err)
		}
	}
	appContainer := &AppContainer{}

	psgConfig := PostgresConfig{
		DB:       "crud",
		User:     "user",
		Password: "pass",
		Port:     nat.Port("5432/tcp"),
	}

	var appCustomConfig = AppConfig{
		Port: nat.Port("8080/tcp"),
	}

	_, err = NewPostgresContainer().Start(ctx, psgConfig, network.Name)
	if err != nil {
		return
	}
	appContainer, err = NewAppContainer().Start(ctx, appCustomConfig, psgConfig, network.Name)

	client := resty.New()
	var token string
	registerUserPayload := domain.RegisterRequestDTO{
		Email:           "msdDaliriyan@gmail.com",
		Password:        "moltafet1",
		ConfirmPassword: "moltafet",
	}

	inCorrectLoginUserPayload := domain.RegisterRequestDTO{
		Email:    "incorrent@incorrect.com",
		Password: "someRandomPassword",
	}

	loginUserPayload := domain.LoginRequestDTO{
		Email:    registerUserPayload.Email,
		Password: registerUserPayload.Password,
	}

	type SendVerificationResponse struct {
		Data int
	}

	var sendVerificationResponse SendVerificationResponse

	baseUrl := appContainer.URL
	Describe("Register", func() {
		It("should register the user successfully", func() {
			resp, _ := client.R().
				SetBody(registerUserPayload).
				Post(fmt.Sprintf("%s/auth/register", baseUrl))
			By("The status code should be 201")
			Expect(resp.StatusCode()).To(Equal(http.StatusCreated))
		})
		It("should not register duplicated error", func() {
			resp, _ := client.R().
				SetBody(registerUserPayload).
				Post(fmt.Sprintf("%s/auth/register", baseUrl))
			By("The status code should be 409")
			Expect(resp.StatusCode()).To(Equal(http.StatusConflict))
		})
	})

	Describe("Email Verification", func() {
		It("Should send verification code successfully", func() {
			resp, _ := client.R().
				SetResult(&sendVerificationResponse).
				Post(fmt.Sprintf("%s/emails/%s/send-verification-code", baseUrl, registerUserPayload.Email))
			Expect(resp.StatusCode()).To(Equal(http.StatusCreated))
		})
		It("Should Verify the created email", func() {
			verifyPayload := &domain.VerifyRequestDTO{
				Code: sendVerificationResponse.Data,
			}
			resp, _ := client.R().
				SetBody(verifyPayload).
				SetResult(&sendVerificationResponse).
				Patch(fmt.Sprintf("%s/emails/%s/verify", baseUrl, registerUserPayload.Email))
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})
	})
	Describe("Login", func() {
		It("should logged-in with created user and get JWT token", func() {
			result := domain.LoginResponseDTO{}
			resp, _ := client.R().
				SetResult(&result).
				SetBody(loginUserPayload).
				Post(fmt.Sprintf("%s/auth/login", baseUrl))
			By("The status code should be 200")
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			token = result.Data.AccessToken
			GinkgoWriter.Println("body", result)
			Expect(result.Data.AccessToken).ToNot(BeEmpty())
		})
		It("should response with 401 status code with incorrect credential", func() {
			resp, _ := client.R().
				SetBody(inCorrectLoginUserPayload).
				Post(fmt.Sprintf("%s/auth/login", baseUrl))
			By("The status code should be 401")
			Expect(resp.StatusCode()).To(Equal(http.StatusUnauthorized))
		})
	})
	Describe("ResetPassword", func() {
		newPassword := "newPassword"
		It("should reset password successfully", func() {
			resetPasswordRequestPayload := domain.ResetPasswordRequestDTO{
				Email:    registerUserPayload.Email,
				Password: newPassword,
			}
			resp, _ := client.R().
				SetBody(resetPasswordRequestPayload).
				Patch(fmt.Sprintf("%s/auth/reset-password", baseUrl))
			By("The status code should be 200")
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})
		It("should not logged-in because password is changed", func() {
			resp, _ := client.R().
				SetBody(loginUserPayload).
				Post(fmt.Sprintf("%s/auth/login", baseUrl))
			By("The status code should be 401")
			Expect(resp.StatusCode()).To(Equal(http.StatusUnauthorized))
		})
	})
	Describe("Users", func() {
		It("should return signed-up user in /me", func() {
			var res = common.ResponseDTO{}

			resp, _ := client.R().
				SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
				SetResult(&res).
				Get(fmt.Sprintf("%s/users/me", baseUrl))
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			By("the user should be same by the registered user")
			email := res.Data.(map[string]interface{})["email"]
			Expect(email).To(Equal(registerUserPayload.Email))
		})
	})
})
