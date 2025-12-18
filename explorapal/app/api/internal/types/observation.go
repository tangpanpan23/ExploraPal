package types

// 观察阶段相关类型定义

type (
	RecognizeImageReq struct {
		ImageUrl string `json:"image_url" desc:"图片URL"`
		Prompt   string `json:"prompt,optional" desc:"分析提示词"`
		Category string `json:"category,optional" desc:"图片类别"`
	}

	RecognizeImageResp struct {
		ObjectName    string   `json:"object_name" desc:"识别出的物体名称"`
		Category      string   `json:"category" desc:"物体类别"`
		Confidence    float32  `json:"confidence" desc:"识别置信度"`
		Description   string   `json:"description" desc:"详细描述"`
		KeyFeatures   []string `json:"key_features" desc:"关键特征"`
		ScientificName string  `json:"scientific_name,optional" desc:"科学名称"`
	}
)
