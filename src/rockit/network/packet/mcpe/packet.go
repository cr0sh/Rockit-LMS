//Package mcpe provides Mojang-defined MCPE packets
package mcpe

import "fmt"

//Packet is a data packet interface, for MCPE Clients
type Packet interface {
	Encode() error
	Decode() error
	GetField(string) interface{}
	SetField(string) interface{}
}

var packetPool map[byte]Packet

//Init initializes packetPool
func Init() {
	packetPool = make(map[byte]Packet)
	registerPacket(LoginPacketHead, *new(LoginPacket))
	registerPacket(PlayStatusPacketHead, *new(PlayStatusPacket))
	registerPacket(DisconnectPacketHead, *new(DisconnectPacket))
	registerPacket(BatchPacketHead, *new(BatchPacket))
	registerPacket(TextPacketHead, *new(TextPacket))
	registerPacket(SetTimePacketHead, *new(SetTimePacket))
	registerPacket(StartGamePacketHead, *new(StartGamePacket))
	registerPacket(AddPlayerPacketHead, *new(AddPlayerPacket))
	registerPacket(RemovePlayerPacketHead, *new(RemovePlayerPacket))
	registerPacket(AddEntityPacketHead, *new(AddEntityPacket))
	registerPacket(RemoveEntityPacketHead, *new(RemoveEntityPacket))
	registerPacket(AddItemEntityPacketHead, *new(AddItemEntityPacket))
	registerPacket(TakeItemEntityPacketHead, *new(TakeItemEntityPacket))
	registerPacket(MoveEntityPacketHead, *new(MoveEntityPacket))
	registerPacket(MovePlayerPacketHead, *new(MovePlayerPacket))
	registerPacket(RemoveBlockPacketHead, *new(RemoveBlockPacket))
	registerPacket(UpdateBlockPacketHead, *new(UpdateBlockPacket))
	registerPacket(AddPaintingPacketHead, *new(AddPaintingPacket))
	registerPacket(ExplodePacketHead, *new(ExplodePacket))
	registerPacket(LevelEventPacketHead, *new(LevelEventPacket))
	registerPacket(TileEventPacketHead, *new(TileEventPacket))
	registerPacket(EntityEventPacketHead, *new(EntityEventPacket))
	registerPacket(MobEffectPacketHead, *new(MobEffectPacket))
	registerPacket(UpdateAttributesPacketHead, *new(UpdateAttributesPacket))
	registerPacket(MobEquipmentPacketHead, *new(MobEquipmentPacket))
	registerPacket(MobArmorEquipmentPacketHead, *new(MobArmorEquipmentPacket))
	registerPacket(InteractPacketHead, *new(InteractPacket))
	registerPacket(UseItemPacketHead, *new(UseItemPacket))
	registerPacket(PlayerActionPacketHead, *new(PlayerActionPacket))
	registerPacket(HurtArmorPacketHead, *new(HurtArmorPacket))
	registerPacket(SetEntityDataPacketHead, *new(SetEntityDataPacket))
	registerPacket(SetEntityMotionPacketHead, *new(SetEntityMotionPacket))
	registerPacket(SetEntityLinkPacketHead, *new(SetEntityLinkPacket))
	registerPacket(SetHealthPacketHead, *new(SetHealthPacket))
	registerPacket(SetSpawnPositionPacketHead, *new(SetSpawnPositionPacket))
	registerPacket(AnimatePacketHead, *new(AnimatePacket))
	registerPacket(RespawnPacketHead, *new(RespawnPacket))
	registerPacket(DropItemPacketHead, *new(DropItemPacket))
	registerPacket(ContainerOpenPacketHead, *new(ContainerOpenPacket))
	registerPacket(ContainerClosePacketHead, *new(ContainerClosePacket))
	registerPacket(ContainerSetSlotPacketHead, *new(ContainerSetSlotPacket))
	registerPacket(ContainerSetDataPacketHead, *new(ContainerSetDataPacket))
	registerPacket(ContainerSetContentPacketHead, *new(ContainerSetContentPacket))
	registerPacket(CraftingDataPacketHead, *new(CraftingDataPacket))
	registerPacket(CraftingEventPacketHead, *new(CraftingEventPacket))
	registerPacket(AdventureSettingsPacketHead, *new(AdventureSettingsPacket))
	registerPacket(TileEntityDataPacketHead, *new(TileEntityDataPacket))
	registerPacket(PlayerInputPacketHead, *new(PlayerInputPacket))
	registerPacket(FullChunkDataPacketHead, *new(FullChunkDataPacket))
	registerPacket(SetDifficultyPacketHead, *new(SetDifficultyPacket))
	registerPacket(ChangeDimensionPacketHead, *new(ChangeDimensionPacket))
	registerPacket(SetPlayerGametypePacketHead, *new(SetPlayerGametypePacket))
	registerPacket(PlayerListPacketHead, *new(PlayerListPacket))
	registerPacket(TelemetryEventPacketHead, *new(TelemetryEventPacket))
}

func registerPacket(head byte, pk Packet) {
	if _, ok := packetPool[head]; ok {
		return
	}
	packetPool[head] = pk
}

//GetPacket gets MCPE Packet from pool with given header, and returns it
func GetPacket(head byte) (pk Packet, err error) {
	if pk, ok := packetPool[head]; !ok {
		return pk, fmt.Errorf("Unimplemented or unregistered packet")
	}
	return pk, nil
}
