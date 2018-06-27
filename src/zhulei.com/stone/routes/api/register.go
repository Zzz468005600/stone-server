package stone

import (
	"errors"
	"github.com/labstack/echo"
	"strings"
	"zhulei.com/stone/context"
	"zhulei.com/stone/db"
)

type registerParams struct {
	Name     string `json:"name"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}

func (p *registerParams) check() error {
	p.Name = strings.TrimSpace(p.Name)
	if p.Name == "" {
		return errors.New("用户名不可为空")
	}
	p.Mobile = strings.TrimSpace(p.Mobile)
	if p.Mobile == "" || len(p.Mobile) != 11 {
		return errors.New("手机号不合法")
	}
	p.Password = strings.TrimSpace(p.Password)
	if p.Password == "" {
		return errors.New("密码不可为空")
	}
	return nil
}

//POST用户注册
var Register = func(c echo.Context) error {
	var p registerParams
	if err := c.Bind(&p); err != nil {
		return context.ErrBadRequest(err.Error())
	}
	if err := p.check(); err != nil {
		return context.ErrBadRequest(err.Error())
	}

	tx, err := db.Pool().Begin()
	if err != nil {
		return context.ErrBadRequest(err.Error())
	}
	defer tx.Rollback()

	var count uint8
	if err := tx.QueryRow(db.StmtUserFindByMobile, p.Mobile).Scan(&count); err != nil {
		return context.ErrBadRequest(err.Error())
	}
	if count != 0 {
		return context.ErrBadRequest("该手机号已被注册")
	}

	tag, err := tx.Exec(db.StmtUserInsert, p.Name, p.Mobile, p.Password)
	if err != nil {
		return context.ErrBadRequest(err.Error())
	}
	if tag.RowsAffected() != 1 {
		return context.ErrBadRequest("注册失败")
	}

	if err := tx.Commit(); err != nil {
		return context.ErrBadRequest(err.Error())
	}

	return context.Say(c, map[string]interface{}{})
}
