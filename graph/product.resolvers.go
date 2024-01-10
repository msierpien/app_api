package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"api/graph/model"
	"context"
	"fmt"
)

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context, id string) (*model.Product, error) {
	panic(fmt.Errorf("not implemented: Product - product"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context, first int, skip *int) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented: Products - products"))
}

// Image is the resolver for the Image field.
func (r *queryResolver) Image(ctx context.Context, id string) (*model.Image, error) {
	panic(fmt.Errorf("not implemented: Image - Image"))
}