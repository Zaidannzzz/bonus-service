package paginate

import (
	"backend-api-test/config"
	paginateModel "backend-api-test/pkg/helpers/databases/paginate/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"sync"
)

var lock = &sync.Mutex{}
var PaginationValueObject *paginateModel.QueryParamPaginationRequest
var defaultPage int = 1       // default value page
var defaultLimit int = 100    //default value limit
var defaultOffset int = 0     //default value offset
var defaultSearch string = "" //default value search
var defaultOrderBy string = "asc"

func PaginationMetadata(
	count *int64,
	limit int,
	page *int,
	endpoint string,
	conf *config.Config,
) paginateModel.ResponseBackPaginationResponse {
	var totalPage float64
	previousPage := *page - 1
	nextPage := *page + 1
	totalData := int(*count)
	totalPage1 := float64(totalData) / float64(limit)
	totalPage1 = math.Ceil(totalPage1)

	if totalPage1 == 0 {
		totalPage = 1
	} else {
		totalPage = totalPage1
	}

	nextPageString := fmt.Sprintf("%s/%spage=%d&limit=%d", conf.Server.Url, endpoint, nextPage, limit)
	previousPageUrlString := fmt.Sprintf("%s/%spage=%d&limit=%d", conf.Server.Url, endpoint, previousPage, limit)
	firstPageUrlString := fmt.Sprintf("%s/%spage=%d&limit=%d", conf.Server.Url, endpoint, 1, limit)
	lastPageUrlString := fmt.Sprintf("%s/%spage=%v&limit=%d", conf.Server.Url, endpoint, totalPage, limit)

	results := paginateModel.ResponseBackPaginationResponse{
		TotalData:        &totalData,
		TotalDataPerPage: &limit,
		CurrentPage:      page,
		PreviousPage:     &previousPage,
		TotalPage:        &totalPage,
		NextPageUrl:      &nextPageString,
		PreviousPageUrl:  &previousPageUrlString,
		FirstPageUrl:     &firstPageUrlString,
		LastPageUrl:      &lastPageUrlString,
	}
	return results
}

func GetPaginationValueObject() *paginateModel.QueryParamPaginationRequest {

	if PaginationValueObject == nil {
		lock.Lock()
		defer lock.Unlock()
		if PaginationValueObject == nil {
			fmt.Println("Initiate Pagination Default Value now.")
			PaginationValueObject = &paginateModel.QueryParamPaginationRequest{
				Page:    func(i int) *int { return &i }(defaultPage),
				Limit:   func(i int) *int { return &i }(defaultLimit),
				Offset:  func(i int) *int { return &i }(defaultOffset),
				Search:  func(i string) *string { return &i }(defaultSearch),
				OrderBy: func(i string) *string { return &i }(defaultOrderBy),
			}
		} else {
			fmt.Println("Pagination instance already created.")
		}
	} else {
		fmt.Println("Pagination instance already created.")
	}

	return PaginationValueObject
}

func QueryParamPaginateTransform(ctx *gin.Context) (*paginateModel.QueryParamPaginationRequest, error) {
	var result paginateModel.QueryParamPaginationRequest
	var resultEntity paginateModel.QueryParamPaginationRequest
	var offset int = *PaginationValueObject.Offset
	var page int = *PaginationValueObject.Page
	var limit int = *PaginationValueObject.Limit
	var search string = *PaginationValueObject.Search
	var orderBy string = *PaginationValueObject.OrderBy

	if err := ctx.ShouldBindQuery(&result); err != nil {
		return nil, err
	}
	if result.Page != nil && *result.Page < 1 {
		return nil, errors.New("Page query params not allowed under 1")
	} else if result.Limit != nil && *result.Limit < 1 {
		return nil, errors.New("Page query params limit not allowed under 1")
	} else {
		// {endpoint}?page=1&limit=1&search=fulan

		if result.OrderBy != nil {
			orderBy = *result.OrderBy
		}

		if result.Limit != nil {
			limit = *result.Limit
		}

		if result.Page != nil {
			page = *result.Page
		}
		var newSearch *string = nil
		if result.Search != nil {
			if len(*result.Search) >= 256 || len(*result.Search) <= 3 {
				return nil, errors.New("length query params search not compatible")
			}
			fmt.Println("search isnt nil")
			search = *result.Search
			newSearch = &search
		}

		if result.Offset != nil {
			offset = ((page - 1) * limit) + *result.Offset
		} else {
			offset = ((page - 1) * limit)
		}

		resultEntity = paginateModel.QueryParamPaginationRequest{
			Page:    &page,
			Offset:  &offset,
			Limit:   &limit,
			Search:  newSearch,
			OrderBy: &orderBy,
		}
	}

	return &resultEntity, nil
}
