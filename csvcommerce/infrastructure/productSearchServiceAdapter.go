package infrastructure

import (
	"context"

	"flamingo.me/flamingo-commerce-adapter-standalone/inMemoryProductSearch/infrastructure"
	productDomain "flamingo.me/flamingo-commerce/product/domain"
	searchDomain "flamingo.me/flamingo-commerce/search/domain"
)

type (
	// ProductSearchServiceAdapter implements methods to search in a product repository
	ProductSearchServiceAdapter struct {
		InMemoryProductRepository *infrastructure.InMemoryProductRepository `inject:""`
	}
)

var (
	_ productDomain.SearchService = &ProductSearchServiceAdapter{}
)

// Search returns a Search Result for the given context and supplied filters
func (p *ProductSearchServiceAdapter) Search(ctx context.Context, filter ...searchDomain.Filter) (productDomain.SearchResult, error) {
	products, err := p.InMemoryProductRepository.Find(filter...)
	if err != nil {
		return productDomain.SearchResult{}, err
	}

	return productDomain.SearchResult{
		Hits: products,
		Result: searchDomain.Result{
			SearchMeta: searchDomain.SearchMeta{
				NumResults: len(products),
			},
		},
	}, nil
}

// SearchBy returns Products prefiltered by the given attribute (also based on additional given Filters) e.g. SearchBy(ctx,"brandCode","apple")
func (p *ProductSearchServiceAdapter) SearchBy(ctx context.Context, attribute string, values []string, filter ...searchDomain.Filter) (productDomain.SearchResult, error) {
	return p.Search(ctx, nil)
}