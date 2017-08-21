package models

import "mhc/manager"

type Name struct {
	Name string
}

func BackupGetAll() manager.Manager {

	var backupList *manager.Manager

	backupList = manager.GetInstance()

	return *backupList
}

func TriggerBackup(name string) manager.Manager {

	var backupList *manager.Manager

	backupList = manager.GetInstance()

	backupList.Name = name
	return *backupList
}