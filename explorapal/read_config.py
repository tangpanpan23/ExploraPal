#!/bin/bash
#
# 配置文件读取工具
# 用于从YAML配置文件中读取视频生成相关的配置信息
#

CONFIG_FILE="${2:-video_generation_config.yaml}"

# 检查配置文件是否存在
if [ ! -f "$CONFIG_FILE" ]; then
    echo "错误: 配置文件 '$CONFIG_FILE' 不存在" >&2
    exit 1
fi

get_config_value() {
    local key="$1"
    local value

    # 使用grep和sed解析简单的YAML结构
    case "$key" in
        "api_key")
            local app_id=$(grep "^  AppID:" "$CONFIG_FILE" | sed 's/.*: *//' | tr -d '"')
            local app_key=$(grep "^  AppKey:" "$CONFIG_FILE" | sed 's/.*: *//' | tr -d '"')
            echo "${app_id}:${app_key}"
            ;;
        "base_url")
            grep "^  BaseURL:" "$CONFIG_FILE" | sed 's/.*: *//' | tr -d '"'
            ;;
        "model")
            grep "^  Model:" "$CONFIG_FILE" | sed 's/.*: *//' | tr -d '"'
            ;;
        "app_id")
            grep "^  AppID:" "$CONFIG_FILE" | sed 's/.*: *//' | tr -d '"'
            ;;
        "app_key")
            grep "^  AppKey:" "$CONFIG_FILE" | sed 's/.*: *//' | tr -d '"'
            ;;
        "duration")
            grep "^  DefaultDuration:" "$CONFIG_FILE" | sed 's/.*: *//' | tr -d '"'
            ;;
        *)
            echo "错误: 不支持的配置项 '$key'" >&2
            echo "支持的配置项: api_key, base_url, model, app_id, app_key, duration" >&2
            exit 1
            ;;
    esac
}

main() {
    if [ $# -lt 1 ]; then
        echo "用法: $0 <配置项> [配置文件]"
        echo "示例:"
        echo "  $0 api_key"
        echo "  $0 base_url"
        echo "  $0 model"
        echo "  $0 app_id video_generation_config.yaml"
        exit 1
    fi

    get_config_value "$1"
}

main "$@"
