package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"
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

// Blocking is the resolver for the blocking field.
func (r *blocksResolver) Blocking(ctx context.Context, obj *models.Blocks) (int, error) {
	panic(fmt.Errorf("not implemented: Blocking - blocking"))
}

// Blocked is the resolver for the blocked field.
func (r *blocksResolver) Blocked(ctx context.Context, obj *models.Blocks) (int, error) {
	panic(fmt.Errorf("not implemented: Blocked - blocked"))
}

// BlockedUser is the resolver for the blocked_user field.
func (r *blocksResolver) BlockedUser(ctx context.Context, obj *models.Blocks) (*models.Users, error) {
	panic(fmt.Errorf("not implemented: BlockedUser - blocked_user"))
}

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input *types.NewAccounts) (*models.Accounts, error) {
	a := &models.Accounts{
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
		Password:    input.Password,
	}
	return a, nil
}

// CreateAccountAndUser is the resolver for the createAccountAndUser field.
func (r *mutationResolver) CreateAccountAndUser(ctx context.Context, account types.NewAccounts, user types.NewUsers) (*models.Users, error) {
	accountInput := &models.Accounts{}
	accountInput.Email = account.Email
	accountInput.Password = account.Password

	userInput := &models.Users{}
	userInput.DisplayName = user.DisplayName

	accountsGraphqlController := controllers.NewAccountsGraphqlController(r.DB, r.Jwt)

	userResponse, res := accountsGraphqlController.Post(ctx, accountInput, userInput)
	return userResponse, res
}

// CreateBlock is the resolver for the createBlock field.
func (r *mutationResolver) CreateBlock(ctx context.Context, input *types.NewBlocks) (*models.Blocks, error) {
	panic(fmt.Errorf("not implemented: CreateBlock - createBlock"))
}

// CreateLike is the resolver for the createLike field.
func (r *mutationResolver) CreateLike(ctx context.Context, input *types.NewLikes) (*models.Likes, error) {
	like := &models.Likes{
		ReceiveUserID: input.ReceiveUserID,
	}
	likesGraphqlController := controllers.NewLikesGraphqlController(r.DB, r.Jwt)

	return likesGraphqlController.Post(ctx, like)
}

// CreateReport is the resolver for the createReport field.
func (r *mutationResolver) CreateReport(ctx context.Context, input *types.NewReports) (*models.Reports, error) {
	panic(fmt.Errorf("not implemented: CreateReport - createReport"))
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*models.Users, error) {
	authorizeGraphqlController := controllers.NewAuthorizeGraphqlController(r.DB, r.Jwt)
	return authorizeGraphqlController.Login(ctx, email, password)
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input *types.NewUsers) (*models.Users, error) {
	panic(fmt.Errorf("not implemented: CreateUsers - createUsers"))
}

// CreateUserSearchFilters is the resolver for the createUserSearchFilters field.
func (r *mutationResolver) CreateUserSearchFilters(ctx context.Context, input *types.NewUserSearchFilters) (*models.UserSearchFilters, error) {
	usf := &models.UserSearchFilters{
		Gender:   input.Gender,
		Location: input.Location,
	}
	userSearchFiltersGraphqlController := controllers.NewUserSearchFilterGraphqlController(r.DB, r.Jwt)

	return userSearchFiltersGraphqlController.Post(ctx, usf)
}

// UpdateUserSearchFilters is the resolver for the updateUserSearchFilters field.
func (r *mutationResolver) UpdateUserSearchFilters(ctx context.Context, input *types.UpdateUserSearchFilters) (*models.UserSearchFilters, error) {
	id, _ := strconv.Atoi(input.ID)
	usf := &models.UserSearchFilters{
		ID:     id,
		Gender: input.Gender,
		// Location: input.Location,
		BodyTypeID:  input.BodyTypeID,
		BloodTypeID: input.BloodTypeID,
	}
	userSearchFiltersGraphqlController := controllers.NewUserSearchFilterGraphqlController(r.DB, r.Jwt)

	return userSearchFiltersGraphqlController.Patch(ctx, usf)
}

// UpdateAccount is the resolver for the updateAccount field.
func (r *mutationResolver) UpdateAccount(ctx context.Context, input *types.UpdateAccounts) (*models.Accounts, error) {
	// account := &models.Accounts{
	// 	ID:          input.ID,
	// 	PhoneNumber: input.PhoneNumber,
	// 	Email:       input.Email,
	// 	Password:    input.Password,
	// }

	accountsGraphqlController := controllers.NewAccountsGraphqlController(r.DB, r.Jwt)

	return accountsGraphqlController.Patch(ctx, input)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input *types.UpdateUsers) (*models.Users, error) {
	// user := &models.Users{
	// 	// ID:          input.ID,
	// 	DisplayName: input.DisplayName,
	// 	// ScreenName:  input.ScreenName,
	// 	Gender:   input.Gender,
	// 	Age:      input.Age,
	// 	Location: input.Location,
	// }

	usersGraphqlController := controllers.NewUsersGraphqlController(r.DB, r.Jwt)

	return usersGraphqlController.Patch(ctx, input)
}

// UpdateUserProfile is the resolver for the updateUserProfile field.
func (r *mutationResolver) UpdateUserProfile(ctx context.Context, input *types.UpdateUserProfiles) (*models.UserProfiles, error) {
	userProfilesGraphqlController := controllers.NewUserProfilesGraphqlController(r.DB)
	return userProfilesGraphqlController.Patch(ctx, input)
}

// CreateVerifyEmail is the resolver for the createVerifyEmail field.
func (r *mutationResolver) CreateVerifyEmail(ctx context.Context, input *types.NewVerifyEmails) (*models.VerifyEmails, error) {
	verifyEmailsGraphqlController := controllers.NewVerifyEmailsGraphqlController(r.DB)

	return verifyEmailsGraphqlController.Post(ctx, input.Email)
}

// Account is the resolver for the account field.
func (r *queryResolver) Account(ctx context.Context, id string) (*models.Accounts, error) {
	accountsGraphqlController := controllers.NewAccountsGraphqlController(r.DB, r.Jwt)

	userID, _ := strconv.Atoi(id)

	return accountsGraphqlController.Get(ctx, userID)
}

// Blocks is the resolver for the blocks field.
func (r *queryResolver) Blocks(ctx context.Context) ([]*models.Blocks, error) {
	panic(fmt.Errorf("not implemented: Blocks - blocks"))
}

// Block is the resolver for the block field.
func (r *queryResolver) Block(ctx context.Context, id string) (*models.Blocks, error) {
	panic(fmt.Errorf("not implemented: Block - block"))
}

// Me is the resolver for the me field.
func (r *queryResolver) Me(ctx context.Context) (*models.ResponseUsers, error) {
	meGraphqlController := controllers.NewMeGraphqlController(r.DB, r.Jwt)

	return meGraphqlController.Get(ctx)
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
func (r *queryResolver) Users(ctx context.Context, first int, after string) (*types.UserConnection, error) {
	usersGraphqlController := controllers.NewUsersGraphqlController(r.DB, r.Jwt)
	return usersGraphqlController.GetList(ctx, first, after)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id string) (*models.ResponseUsers, error) {
	ids, _ := strconv.Atoi(id)
	usersGraphqlController := controllers.NewUsersGraphqlController(r.DB, r.Jwt)

	return usersGraphqlController.Get(ctx, ids)
}

// UserSearchFilters is the resolver for the user_search_filters field.
func (r *queryResolver) UserSearchFilters(ctx context.Context) (*models.UserSearchFilters, error) {
	userSearchFiltersGraphqlController := controllers.NewUserSearchFilterGraphqlController(r.DB, r.Jwt)

	return userSearchFiltersGraphqlController.Get(ctx)
}

// VerifyEmail is the resolver for the verify_email field.
func (r *queryResolver) VerifyEmail(ctx context.Context, code string) (*models.VerifyEmails, error) {
	verifyEmailsGraphqlController := controllers.NewVerifyEmailsGraphqlController(r.DB)

	return verifyEmailsGraphqlController.Get(ctx, code)
}

// HeightID is the resolver for the height_id field.
func (r *userProfilesResolver) HeightID(ctx context.Context, obj *models.UserProfiles) (int, error) {
	panic(fmt.Errorf("not implemented: HeightID - height_id"))
}

// ResidenceID is the resolver for the residence_id field.
func (r *userProfilesResolver) ResidenceID(ctx context.Context, obj *models.UserProfiles) (int, error) {
	panic(fmt.Errorf("not implemented: ResidenceID - residence_id"))
}

// HometownID is the resolver for the hometown_id field.
func (r *userProfilesResolver) HometownID(ctx context.Context, obj *models.UserProfiles) (int, error) {
	panic(fmt.Errorf("not implemented: HometownID - hometown_id"))
}

// IndealFirstEncointerID is the resolver for the indeal_first_encointer_id field.
func (r *userProfilesResolver) IndealFirstEncointerID(ctx context.Context, obj *models.UserProfiles) (int, error) {
	panic(fmt.Errorf("not implemented: IndealFirstEncointerID - indeal_first_encointer_id"))
}

// PresonalityTypeID is the resolver for the presonality_type_id field.
func (r *userProfilesResolver) PresonalityTypeID(ctx context.Context, obj *models.UserProfiles) (int, error) {
	panic(fmt.Errorf("not implemented: PresonalityTypeID - presonality_type_id"))
}

// IsVerifiedEmail is the resolver for the is_verified_email field.
func (r *usersResolver) IsVerifiedEmail(ctx context.Context, obj *models.Users) (bool, error) {
	panic(fmt.Errorf("not implemented: IsVerifiedEmail - is_verified_email"))
}

// Accounts returns AccountsResolver implementation.
func (r *Resolver) Accounts() AccountsResolver { return &accountsResolver{r} }

// Blocks returns BlocksResolver implementation.
func (r *Resolver) Blocks() BlocksResolver { return &blocksResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// UserProfiles returns UserProfilesResolver implementation.
func (r *Resolver) UserProfiles() UserProfilesResolver { return &userProfilesResolver{r} }

// Users returns UsersResolver implementation.
func (r *Resolver) Users() UsersResolver { return &usersResolver{r} }

type accountsResolver struct{ *Resolver }
type blocksResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userProfilesResolver struct{ *Resolver }
type usersResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *likesResolver) SendUserID(ctx context.Context, obj *models.Likes) (int, error) {
	panic(fmt.Errorf("not implemented: SendUserID - send_user_id"))
}
func (r *likesResolver) ReceiveUserID(ctx context.Context, obj *models.Likes) (int, error) {
	panic(fmt.Errorf("not implemented: ReceiveUserID - receive_user_id"))
}

type likesResolver struct{ *Resolver }

func (r *usersResolver) ScreenName(ctx context.Context, obj *models.Users) (string, error) {
	panic(fmt.Errorf("not implemented: ScreenName - screen_name"))
}
