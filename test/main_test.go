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
		Data string
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
				UserID:   1,
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
		It("should return signed-up user in authenticated user", func() {
			var res = common.ResponseDTO{}
			resp, _ := client.R().
				SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
				SetResult(&res).
				Get(fmt.Sprintf("%s/user", baseUrl))
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
			By("the user should be same by the registered user")
			email := res.Data.(map[string]interface{})["email"]
			Expect(email).To(Equal(registerUserPayload.Email))
		})
	})
	Describe("Assets", func() {
		UserId := "1"
		p := domain.StoreTransactionRequestDTO{
			Amount:            10.2,
			Description:       "description...",
			TransactionTypeId: 1,
		}

		It("should can deposit", func() {
			resp, _ := client.R().
				SetBody(p).
				SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
				Post(fmt.Sprintf("%s/users/%s/transactions", baseUrl, UserId))
			By("The status code should be 200")
			Expect(resp.StatusCode()).To(Equal(http.StatusCreated))
		})
		It("should the amount of the asset increased after depositing", func() {
			r := common.ResponseDTO{}
			resp, _ := client.R().
				SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
				SetResult(&r).
				Get(fmt.Sprintf("%s/user/asset", baseUrl))
			By("The status code should be 200")
			Expect(resp.StatusCode()).To(Equal(200))
		})
	})
	Describe("Plans", func() {
		It("should be able to active a plan", func() {

		})
		It("should be able to inactive a plan", func() {

		})
		It("should one plan exist at least", func() {
			r := common.ResponseDTO{}
			resp, _ := client.R().
				SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
				SetResult(&r).
				Get(fmt.Sprintf("%s/plans/", baseUrl))

			Expect(resp.StatusCode()).To(Equal(200))
			Expect(r.Data).To(Not(BeEmpty()))

		})
		It("should create plan successfully", func() {
		})
	})
	Describe("Orders", func() {

		It("should be able to create order", func() {
			po := domain.StoreForAuthenticateUserDTO{
				PlanID: 1,
			}
			resp, _ := client.R().
				SetBody(po).
				SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
				Post(fmt.Sprintf("%s/user/orders", baseUrl))
			By("The status code should be 201")
			Expect(resp.StatusCode()).To(Equal(http.StatusCreated))
		})
		It("should be able to accept order", func() {
			p := domain.UpdateOrderDTO{OrderStatusId: 2}
			resp, _ := client.R().
				SetBody(p).
				SetHeader("Authorization", fmt.Sprintf("Bearer %s", token)).
				Patch(fmt.Sprintf("%s/orders/%s", baseUrl, "1"))
			By("The status code should be 200")
			Expect(resp.StatusCode()).To(Equal(http.StatusOK))
		})
	})
	Describe("Transactions plan-request", func() {
		It("should be able to create settled transaction plan-request", func() {})
		It("should be accept transaction plan-request", func() {})
	})
	Describe("User's Plan", func() {
		It("should be able to get list of registered user's plans ", func() {
		})
		It("should be able to deposit", func() {
		})
		It("should be able to register user in one plan", func() {})
	})
})
