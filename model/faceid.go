package model

type FaceID struct {
	ID         string                 `json:"id" `          // ID
	SourceType string                 `json:"source_type" ` // 资源类型
	SourceHash string                 `json:"source_hash" ` // 资源hash
	Algorithm  string                 `json:"algorithm" `   // hash 算法
	Labels     []string               `json:"labels" `      // 标签
	Metadata   map[string]interface{} `json:"metadata" `    // 元数据
	Timestamp  int64                  `json:"timestamp"`    // 时间戳(s)
}

type HistoryFaceIDs struct {
	ID         string                 `json:"id" `          // ID
	StartTime  int64                  `json:"start_time" `  // 查询开始时间
	EndTime    int64                  `json:"end_time" `    // 查询结束时间
	Labels     []string               `json:"labels" `      // 标签
	Timestamp  int64                  `json:"timestamp"`    // 时间戳(s)
}

//type ContractList struct {
//	ID         string                 `json:"id" `          // ID
//	NetworkD   string                 `json:"network_id" `  // 网络类型
//	Channel    string                 `json:"channel" `     // 通道
//	Timestamp  int64                  `json:"timestamp"`    // 时间戳(s)
//}