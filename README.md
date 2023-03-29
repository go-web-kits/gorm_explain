# GORM Explain

Auto SQL Explanation  
性能分析必备 -- 自动跑 SQL 的 Explain（基于 GORM 回调）

Maintainers: @will.huang  
测试覆盖率：   
状态：可用

## Setup

```go
gorm_explain.Register(dbx.Conn().DB)
```

## Usage

只需要设置一个环境变量即可，如下：
```bash
$ EXPLAIN=true go test
```

效果：
