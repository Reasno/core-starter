### otgorm

#### 注册

> config/app.go

```go
package config

import (
    "github.com/DoNewsCode/core/di"
    "github.com/DoNewsCode/core/otgorm"
    // ... others
)

var (
    // Register providers
    Providers = []di.Deps{
        // ... others
        otgorm.Providers(),
    }

    // Register modules
    Modules = []interface{}{
        // ... others
        otgorm.New,
    }
)
```

#### 使用

- 数据库迁移与数据填充

在 `app/Module` 里实现 `ProvideMigration()` 方法, 然后执行

> app/module.go

```
func (m Module) ProvideMigration() []*otgorm.Migration {
    return entities.Migrations()
}
```

> app/entities/migrations.go

```
func Migrations() []*otgorm.Migration {
	return []*otgorm.Migration{
		//{
		//	ID: "202103042231", // 推荐使用日期 格式为 YYYYMMDDHHii
		//	Migrate: func(db *gorm.DB) error {
		//		type User struct {
		//			gorm.Model
		//			Email         string    `gorm:"size:255;uniqueIndex;not null"`
		//			Username      string    `gorm:"size:255;uniqueIndex;not null"`
		//			Password      string    `gorm:"size:255"`
		//			EmailVerifyAt time.Time `gorm:"email_verify_at"`
		//			RememberToken string    `gorm:"remember_token"`
		//		}
		//		return db.AutoMigrate(&User{})
		//	},
		//	Rollback: func(db *gorm.DB) error {
		//		type User struct{}
		//		return db.Migrator().DropTable(&User{})
		//	},
		//},
	}
}
```
