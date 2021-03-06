package repos

import (
	"code.db.cafe/wombot/internal/database"
	"code.db.cafe/wombot/internal/database/entities"
)

var Guild = database.NewRepository[entities.Guild, string](database.Db, entities.Guild{})
var User = database.NewRepository[entities.User, string](database.Db, entities.User{})
