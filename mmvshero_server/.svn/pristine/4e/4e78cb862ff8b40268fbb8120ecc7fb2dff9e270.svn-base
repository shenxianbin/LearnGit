package logic

import (
	. "galaxy"
)

type Manager struct {
	role_sid_list map[int64]IRole
	role_uid_list map[int64]IRole
}

var m *Manager

func Init() {
	m = new(Manager)
	m.role_sid_list = make(map[int64]IRole)
	m.role_uid_list = make(map[int64]IRole)
}

func AddRoleBySid(sid int64, role IRole) {
	m.role_sid_list[sid] = role
}

func AddRoleByUid(uid int64, role IRole) {
	m.role_uid_list[uid] = role
	LogDebug("Now OnLine Role : ", len(m.role_uid_list))
}

func RemRoleBySid(sid int64) {
	delete(m.role_sid_list, sid)
}

func RemRoleByUid(uid int64) {
	delete(m.role_uid_list, uid)
}

func RemoveAll() {
	m.role_sid_list = make(map[int64]IRole)
	m.role_uid_list = make(map[int64]IRole)
}

func GetRoleBySid(sid int64) IRole {
	if role, has := m.role_sid_list[sid]; has {
		return role
	}
	return nil
}

func GetRoleByUid(uid int64) IRole {
	if role, has := m.role_uid_list[uid]; has {
		return role
	}
	return nil
}

func GetAllRoleBySid() map[int64]IRole {
	return m.role_sid_list
}

func GetAllRoleByUid() map[int64]IRole {
	return m.role_uid_list
}
