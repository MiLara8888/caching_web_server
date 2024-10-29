package documentservice

import (
	"encoding/json"
	"io"

	er "github.com/MiLara8888/caching_web_server/pkg/errors"
	"github.com/MiLara8888/caching_web_server/pkg/storage"
	"github.com/gin-gonic/gin"
)

func (m *Rest) RegisterUser(c *gin.Context) {

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(er.StatusBadRequest.Code(), storage.ErrSerializer{Error: *er.StatusBadRequest})
		return
	}

	data := &storage.RegisterSerializer{}

	err = json.Unmarshal(jsonData, &data)
	if err != nil || (!data.Valid()){
		c.AbortWithStatusJSON(er.StatusBadRequest.Code(), storage.ErrSerializer{Error: *er.StatusBadRequest})
		return
	}

	//проверка токена админа
	if m.TokenAdmin!=data.TokenAdmin{
		c.AbortWithStatusJSON(er.StatusForbidden.Code(), storage.ErrSerializer{Error: *er.StatusForbidden})
		return
	}

}
