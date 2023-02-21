package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
	"github.com/takeuchi-shogo/k8s-go-sample/graphql/types"
	"github.com/takeuchi-shogo/k8s-go-sample/interface/controllers"
)

// LoginStatus is the resolver for the login_status field.
func (r *accountsResolver) LoginStatus(ctx context.Context, obj *models.Accounts) (bool, error) {
	panic(fmt.Errorf("not implemented: LoginStatus - login_status"))
}

// AccessLevel is the resolver for the access_level field.
func (r *accountsResolver) AccessLevel(ctx context.Context, obj *models.Accounts) (string, error) {
	panic(fmt.Errorf("not implemented: AccessLevel - access_level"))
}

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input *types.NewAccounts) (*models.Accounts, error) {
	a := &models.Accounts{
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
		Password:    input.Password,
	}
	log.Println(a)
	return a, nil
}

// CreateAccountAndUser is the resolver for the createAccountAndUser field.
func (r *mutationResolver) CreateAccountAndUser(ctx context.Context, account types.NewAccounts, user types.NewUsers) (*models.Users, error) {
	accountInput := &models.Accounts{}
	accountInput.Email = account.Email
	accountInput.Password = account.Password

	userInput := &models.Users{}
	userInput.DisplayName = user.DisplayName

	accountsGraphqlController := controllers.NewAccountsGraphqlController(r.DB)

	userResponse, res := accountsGraphqlController.Post(ctx, accountInput, userInput)
	return userResponse, res
}

// CreateBlock is the resolver for the createBlock field.
func (r *mutationResolver) CreateBlock(ctx context.Context, input *types.NewBlocks) (*models.Blocks, error) {
	panic(fmt.Errorf("not implemented: CreateBlock - createBlock"))
}

// CreateReport is the resolver for the createReport field.
func (r *mutationResolver) CreateReport(ctx context.Context, input *types.NewReports) (*models.Reports, error) {
	panic(fmt.Errorf("not implemented: CreateReport - createReport"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input *types.NewLogin) (*models.Users, error) {
	fmt.Println("koko")
	authorizeGraphqlController := controllers.NewAuthorizeGraphqlController(r.DB, r.Jwt)
	return &models.Users{}, authorizeGraphqlController.Login(ctx, input.Email, input.Password)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *types.NewUsers) (*models.Users, error) {
	panic(fmt.Errorf("not implemented: CreateUsers - createUsers"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input *types.UpdateUsers) (*models.Users, error) {
	user := &models.Users{
		ID:          input.ID,
		DisplayName: input.DisplayName,
		ScreenName:  input.ScreenName,
		Gender:      input.Gender,
		Location:    input.Location,
	}

	usersGraphqlController := controllers.NewUsersGraphqlController(r.DB)

	return usersGraphqlController.Patch(ctx, user)
}

// CreateVerifyEmail is the resolver for the createVerifyEmail field.
func (r *mutationResolver) CreateVerifyEmail(ctx context.Context, input *types.NewVerifyEmails) (*models.VerifyEmails, error) {
	verifyEmailsGraphqlController := controllers.NewVerifyEmailsGraphqlController(r.DB)

	return verifyEmailsGraphqlController.Post(ctx, input.Email)
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, id string) (*models.Accounts, error) {
	panic(fmt.Errorf("not implemented: Account - account"))
}

// Blocks is the resolver for the blocks field.
func (r *queryResolver) Blocks(ctx context.Context) ([]*models.Blocks, error) {
	panic(fmt.Errorf("not implemented: Blocks - blocks"))
}

// Block is the resolver for the block field.
func (r *queryResolver) Block(ctx context.Context, id string) (*models.Blocks, error) {
	panic(fmt.Errorf("not implemented: Block - block"))
}

// Reports is the resolver for the reports field.
func (r *queryResolver) Reports(ctx context.Context) ([]*models.Reports, error) {
	panic(fmt.Errorf("not implemented: Reports - reports"))
}

// Report is the resolver for the report field.
func (r *queryResolver) Report(ctx context.Context, id string) (*models.Reports, error) {
	panic(fmt.Errorf("not implemented: Report - report"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.Users, error) {
	usersGraphqlController := controllers.NewUsersGraphqlController(r.DB)
	fmt.Println(ctx.Value("users"))
	return usersGraphqlController.GetList(ctx)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.Users, error) {
	// pass, _ := utils.GenerateFromPassword("pass")
	ids, _ := strconv.Atoi(id)
	r.user = &models.Users{
		ID: ids,
		// Name:     "test" + id,
		// Password: pass,
	}
	return r.user, nil
}

// VerifyEmail is the resolver for the verify_email field.
func (r *queryResolver) VerifyEmail(ctx context.Context, code string) (*models.VerifyEmails, error) {
	verifyEmailsGraphqlController := controllers.NewVerifyEmailsGraphqlController(r.DB)

	return verifyEmailsGraphqlController.Get(ctx, code)
}

// IsVerifiedEmail is the resolver for the is_verified_email field.
func (r *usersResolver) IsVerifiedEmail(ctx context.Context, obj *models.Users) (bool, error) {
	panic(fmt.Errorf("not implemented: IsVerifiedEmail - is_verified_email"))
}

// Accounts returns AccountsResolver implementation.
func (r *Resolver) Accounts() AccountsResolver { return &accountsResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Users returns UsersResolver implementation.
func (r *Resolver) Users() UsersResolver { return &usersResolver{r} }

type accountsResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type usersResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *usersResolver) ScreenName(ctx context.Context, obj *models.Users) (string, error) {
	panic(fmt.Errorf("not implemented: ScreenName - screen_name"))
}
