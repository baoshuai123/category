package handler

import (
	"context"
	"taobao/category/common"
	"taobao/category/domain/model"
	"taobao/category/domain/service"
	category "taobao/category/proto/category"

	log "github.com/micro/go-micro/v2/logger"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

//提供创建分类的服务
func (c *Category) CrateCategory(ctx context.Context, request *category.CategoryRequest, response *category.CrateCategoryResponse) error {
	categoryData := &model.Category{}
	//赋值
	err := common.SwapTo(request, categoryData)
	if err != nil {
		return err
	}
	categoryId, err := c.CategoryDataService.AddCategory(categoryData)
	if err != nil {
		return err
	}
	response.Message = "分类添加成功"
	response.CategoryId = categoryId
	return nil
}

//提供分类更新的服务
func (c *Category) UpdateCategory(ctx context.Context, request *category.CategoryRequest, response *category.UpdateCategoryResponse) error {
	categoryData := &model.Category{}
	err := common.SwapTo(request, categoryData)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(categoryData)
	if err != nil {
		return err
	}
	response.Message = "分类更新成功"
	return nil
}

//提供分类删除的服务
func (c *Category) DeleteCategory(ctx context.Context, request *category.DeleteCategoryRequest, response *category.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(request.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "删除成功"
	return nil
}

//根据分类名称查找分类
func (c *Category) FindCategoryByName(ctx context.Context, request *category.FindByNameRequest, response *category.CategoryResponse) error {
	categoryData, err := c.CategoryDataService.FindCategoryByName(request.CategoryName)
	if err != nil {
		return err
	}
	return common.SwapTo(categoryData, response)
}

//根据分类ID查找分类
func (c *Category) FindCategoryByID(ctx context.Context, request *category.FindByIDRequest, response *category.CategoryResponse) error {
	categoryData, err := c.CategoryDataService.FindCategoryByID(request.CategoryId)
	if err != nil {
		return err
	}

	return common.SwapTo(categoryData, response)
}

//查找层级下所有分类
func (c *Category) FindCategoryByLevel(ctx context.Context, request *category.FindByLevelRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByLevel(request.Level)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)
	return nil
}

// 给返回的切片赋值
func categoryToResponse(categorySlice []model.Category, response *category.FindAllResponse) {
	for _, categoryData := range categorySlice {
		cr := &category.CategoryResponse{}
		err := common.SwapTo(categoryData, cr)
		if err != nil {
			log.Error(err)
			break
		}
		response.Category = append(response.Category, cr)
	}
}

func (c *Category) FindCategoryByParent(ctx context.Context, request *category.FindByParentRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindCategoryByParent(request.ParentId)
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)
	return nil
}

func (c *Category) FindAllCategory(ctx context.Context, request *category.FindAllRequest, response *category.FindAllResponse) error {
	categorySlice, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	categoryToResponse(categorySlice, response)
	return nil
}
