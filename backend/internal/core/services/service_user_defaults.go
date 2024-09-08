package services

import (
	"github.com/hay-kot/homebox/backend/internal/data/repo"
)

func defaultLocations() []repo.LocationCreate {
	return []repo.LocationCreate{
		{
			Name: "客厅",
		},
		{
			Name: "车库",
		},
		{
			Name: "厨房",
		},
		{
			Name: "卧室",
		},
		{
			Name: "浴室",
		},
		{
			Name: "办公室",
		},
		{
			Name: "阁楼",
		},
		{
			Name: "地下室",
		},
	}
}

func defaultLabels() []repo.LabelCreate {
	return []repo.LabelCreate{
		{
			Name: "家用电器",
		},
		{
			Name: "物联网",
		},
		{
			Name: "电子产品",
		},
		{
			Name: "服务器",
		},
		{
			Name: "常规",
		},
		{
			Name: "重要",
		},
	}
}
