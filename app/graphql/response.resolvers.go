package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"fmt"

	"github.com/takeuchi-shogo/k8s-go-sample/domain/models"
)

// Purpose is the resolver for the purpose field.
func (r *responseUserProfilesResolver) Purpose(ctx context.Context, obj *models.ResponseUserProfiles) (*string, error) {
	panic(fmt.Errorf("not implemented: Purpose - purpose"))
}

// ResponseUserProfiles returns ResponseUserProfilesResolver implementation.
func (r *Resolver) ResponseUserProfiles() ResponseUserProfilesResolver {
	return &responseUserProfilesResolver{r}
}

type responseUserProfilesResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *responseUserProfilesResolver) PresonalityTypeID(ctx context.Context, obj *models.ResponseUserProfiles) (int, error) {
	panic(fmt.Errorf("not implemented: PresonalityTypeID - presonality_type_id"))
}
func (r *responseUserProfilesResolver) PresonalityType(ctx context.Context, obj *models.ResponseUserProfiles) (string, error) {
	panic(fmt.Errorf("not implemented: PresonalityType - presonality_type"))
}
func (r *responseUserProfilesResolver) IndealFirstEncointerID(ctx context.Context, obj *models.ResponseUserProfiles) (int, error) {
	panic(fmt.Errorf("not implemented: IndealFirstEncointerID - indeal_first_encointer_id"))
}
func (r *responseUserProfilesResolver) IndealFirstEncointer(ctx context.Context, obj *models.ResponseUserProfiles) (string, error) {
	panic(fmt.Errorf("not implemented: IndealFirstEncointer - indeal_first_encointer"))
}
func (r *responseUserProfilesResolver) EducationalID(ctx context.Context, obj *models.ResponseUserProfiles) (int, error) {
	panic(fmt.Errorf("not implemented: EducationalID - educational_id"))
}
func (r *responseUserProfilesResolver) Educational(ctx context.Context, obj *models.ResponseUserProfiles) (string, error) {
	panic(fmt.Errorf("not implemented: Educational - educational"))
}
func (r *responseUserProfilesResolver) JobTitleID(ctx context.Context, obj *models.ResponseUserProfiles) (int, error) {
	panic(fmt.Errorf("not implemented: JobTitleID - job_title_id"))
}
