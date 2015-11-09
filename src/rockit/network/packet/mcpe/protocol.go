package mcpe

const (
	//Protocol defines current MCPE protocol version
	Protocol = 34
	//LoginPacketHead is a header constant for LoginPacket
	LoginPacketHead = 0x8f
	//PlayStatusPacketHead is a header constant for PlayStatusPacket
	PlayStatusPacketHead = 0x90
	//DisconnectPacketHead is a header constant for DisconnectPacket
	DisconnectPacketHead = 0x91
	//BatchPacketHead is a header constant for BatchPacket
	BatchPacketHead = 0x92
	//TextPacketHead is a header constant for TextPacket
	TextPacketHead = 0x93
	//SetTimePacketHead is a header constant for SetTimePacket
	SetTimePacketHead = 0x94
	//StartGamePacketHead is a header constant for StartGamePacket
	StartGamePacketHead = 0x95
	//AddPlayerPacketHead is a header constant for AddPlayerPacket
	AddPlayerPacketHead = 0x96
	//RemovePlayerPacketHead is a header constant for RemovePlayerPacket
	RemovePlayerPacketHead = 0x97
	//AddEntityPacketHead is a header constant for AddEntityPacket
	AddEntityPacketHead = 0x98
	//RemoveEntityPacketHead is a header constant for RemoveEntityPacket
	RemoveEntityPacketHead = 0x99
	//AddItemEntityPacketHead is a header constant for AddItemEntityPacket
	AddItemEntityPacketHead = 0x9a
	//TakeItemEntityPacketHead is a header constant for TakeItemEntityPacket
	TakeItemEntityPacketHead = 0x9b
	//MoveEntityPacketHead is a header constant for MoveEntityPacket
	MoveEntityPacketHead = 0x9c
	//MovePlayerPacketHead is a header constant for MovePlayerPacket
	MovePlayerPacketHead = 0x9d
	//RemoveBlockPacketHead is a header constant for RemoveBlockPacket
	RemoveBlockPacketHead = 0x9e
	//UpdateBlockPacketHead is a header constant for UpdateBlockPacket
	UpdateBlockPacketHead = 0x9f
	//AddPaintingPacketHead is a header constant for AddPaintingPacket
	AddPaintingPacketHead = 0xa0
	//ExplodePacketHead is a header constant for ExplodePacket
	ExplodePacketHead = 0xa1
	//LevelEventPacketHead is a header constant for LevelEventPacket
	LevelEventPacketHead = 0xa2
	//TileEventPacketHead is a header constant for TileEventPacket
	TileEventPacketHead = 0xa3
	//EntityEventPacketHead is a header constant for EntityEventPacket
	EntityEventPacketHead = 0xa4
	//MobEffectPacketHead is a header constant for MobEffectPacket
	MobEffectPacketHead = 0xa5
	//UpdateAttributesPacketHead is a header constant for UpdateAttributesPacket
	UpdateAttributesPacketHead = 0xa6
	//MobEquipmentPacketHead is a header constant for MobEquipmentPacket
	MobEquipmentPacketHead = 0xa7
	//MobArmorEquipmentPacketHead is a header constant for MobArmorEquipmentPacket
	MobArmorEquipmentPacketHead = 0xa8
	//InteractPacketHead is a header constant for InteractPacket
	InteractPacketHead = 0xa9
	//UseItemPacketHead is a header constant for UseItemPacket
	UseItemPacketHead = 0xaa
	//PlayerActionPacketHead is a header constant for PlayerActionPacket
	PlayerActionPacketHead = 0xab
	//HurtArmorPacketHead is a header constant for HurtArmorPacket
	HurtArmorPacketHead = 0xac
	//SetEntityDataPacketHead is a header constant for SetEntityDataPacket
	SetEntityDataPacketHead = 0xad
	//SetEntityMotionPacketHead is a header constant for SetEntityMotionPacket
	SetEntityMotionPacketHead = 0xae
	//SetEntityLinkPacketHead is a header constant for SetEntityLinkPacket
	SetEntityLinkPacketHead = 0xaf
	//SetHealthPacketHead is a header constant for SetHealthPacket
	SetHealthPacketHead = 0xb0
	//SetSpawnPositionPacketHead is a header constant for SetSpawnPositionPacket
	SetSpawnPositionPacketHead = 0xb1
	//AnimatePacketHead is a header constant for AnimatePacket
	AnimatePacketHead = 0xb2
	//RespawnPacketHead is a header constant for RespawnPacket
	RespawnPacketHead = 0xb3
	//DropItemPacketHead is a header constant for DropItemPacket
	DropItemPacketHead = 0xb4
	//ContainerOpenPacketHead is a header constant for ContainerOpenPacket
	ContainerOpenPacketHead = 0xb5
	//ContainerClosePacketHead is a header constant for ContainerClosePacket
	ContainerClosePacketHead = 0xb6
	//ContainerSetSlotPacketHead is a header constant for ContainerSetSlotPacket
	ContainerSetSlotPacketHead = 0xb7
	//ContainerSetDataPacketHead is a header constant for ContainerSetDataPacket
	ContainerSetDataPacketHead = 0xb8
	//ContainerSetContentPacketHead is a header constant for ContainerSetContentPacket
	ContainerSetContentPacketHead = 0xb9
	//CraftingDataPacketHead is a header constant for CraftingDataPacket
	CraftingDataPacketHead = 0xba
	//CraftingEventPacketHead is a header constant for CraftingEventPacket
	CraftingEventPacketHead = 0xbb
	//AdventureSettingsPacketHead is a header constant for AdventureSettingsPacket
	AdventureSettingsPacketHead = 0xbc
	//TileEntityDataPacketHead is a header constant for TileEntityDataPacket
	TileEntityDataPacketHead = 0xbd
	//PlayerInputPacketHead is a header constant for PlayerInputPacket
	PlayerInputPacketHead = 0xbe
	//FullChunkDataPacketHead is a header constant for FullChunkDataPacket
	FullChunkDataPacketHead = 0xbf
	//SetDifficultyPacketHead is a header constant for SetDifficultyPacket
	SetDifficultyPacketHead = 0xc0
	//ChangeDimensionPacketHead is a header constant for ChangeDimensionPacket
	ChangeDimensionPacketHead = 0xc1
	//SetPlayerGametypePacketHead is a header constant for SetPlayerGametypePacket
	SetPlayerGametypePacketHead = 0xc2
	//PlayerListPacketHead is a header constant for PlayerListPacket
	PlayerListPacketHead = 0xc3
	//TelemetryEventPacketHead is a header constant for TelemetryEventPacket
	TelemetryEventPacketHead = 0xc4
)
