package proto

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type SimplePacket struct {
	PacketHead uint32
	PacketType uint32
}

type ACKPacket struct {
	SimplePacket
	PacketTrail uint32
}

type ProfileBusyPacket struct {
	SimplePacket
	Unk0 uint32
}
type ByePacket struct {
	SimplePacket
	ProfileID uint32
}

type EnterProfilePacket struct {
	SimplePacket
	PlayerID  uint32
	MachineID uint32
	ProfileID uint32
}

type RequestWorldBestPacket struct {
	SimplePacket
}

type RequestRankModePacket struct {
	SimplePacket
}

type KeepAlivePacket struct {
	SimplePacket
	PacketTrail uint32
}

type MachineInfoPacket struct {
	PacketHead     uint32        //    0x00 0x0000001
	PacketType     uint32        //    0x04 0x1000011
	MachineID      uint32        //    0x08
	DongleID       uint32        //    0x0C
	CountryID      uint32        //    0x10
	MacAddress     PIUMacAddress //    0x14
	Version        PIUString12   //    0x28
	Processor      PIUString128  //    0x34
	MotherBoard    PIUString128  //    0xb4
	GraphicsCard   PIUString128  //    0x134
	HDDSerial      PIUString32   //    0x1b4
	USBMode        PIUString128  //    Mode 1.0 / Mode 1.1 / Mode 2.0
	Memory         uint32        //    0x254
	ConfigMagic    uint32        //    0x258
	Unk3           uint32        //    0x25c
	Unk4           uint32        //    0x260
	Unk5           uint32        //    0x264
	Unk6           uint32        //    0x268
	Unk7           uint32        //    0x26c
	Unk8           uint32        //    0x270
	Unk9           uint32        //    0x274
	Unk10          uint32        //    0x278
	Unk11          uint32        //    0x27c
	Unk12          uint32
	Unk13          uint32
	Unk14          uint32
	Unk15          uint32
	Unk16          uint32
	Unk17          uint32
	Unk18          uint32
	Unk19          [76]uint8
	NetworkAddress PIUString16
}

func (p *MachineInfoPacket) String() string {
	s := "Machine Packet\n"

	s += fmt.Sprintf("\tPacketHead: %d (0x%x)\n", p.PacketHead, p.PacketHead)
	s += fmt.Sprintf("\tPacketType: %d (0x%x)\n", p.PacketType, p.PacketType)
	s += fmt.Sprintf("\tMachineID: %d (0x%x)\n", p.MachineID, p.MachineID)
	s += fmt.Sprintf("\tDongle ID: %d (0x%x)\n", p.DongleID, p.DongleID)
	s += fmt.Sprintf("\tCountry ID: %d (0x%x)\n", p.CountryID, p.CountryID)
	s += fmt.Sprintf("\tMac Address: %s\n", p.MacAddress)
	s += fmt.Sprintf("\tVersion: %s\n", p.Version)
	s += fmt.Sprintf("\tProcessor: %s\n", p.Processor)
	s += fmt.Sprintf("\tMother Board: %s\n", p.MotherBoard)
	s += fmt.Sprintf("\tGraphics Card: %s\n", p.GraphicsCard)
	s += fmt.Sprintf("\tHDD Serial: %s\n", p.HDDSerial)
	s += fmt.Sprintf("\tUSB Mode: %s\n", p.USBMode)
	s += fmt.Sprintf("\tMemory: %d\n", p.Memory)
	s += fmt.Sprintf("\tConfig Magic: %d (0x%x)\n", p.ConfigMagic, p.ConfigMagic)
	s += fmt.Sprintf("\tNet Address: %s\n", p.NetworkAddress)

	s += fmt.Sprintf("\tUnknown uint32_t  3: %d (0x%x)\n", p.Unk3, p.Unk3)
	s += fmt.Sprintf("\tUnknown uint32_t  4: %d (0x%x)\n", p.Unk4, p.Unk4)
	s += fmt.Sprintf("\tUnknown uint32_t  5: %d (0x%x)\n", p.Unk5, p.Unk5)
	s += fmt.Sprintf("\tUnknown uint32_t  6: %d (0x%x)\n", p.Unk6, p.Unk6)
	s += fmt.Sprintf("\tUnknown uint32_t  7: %d (0x%x)\n", p.Unk7, p.Unk7)
	s += fmt.Sprintf("\tUnknown uint32_t  8: %d (0x%x)\n", p.Unk8, p.Unk8)
	s += fmt.Sprintf("\tUnknown uint32_t  9: %d (0x%x)\n", p.Unk9, p.Unk9)
	s += fmt.Sprintf("\tUnknown uint32_t 10: %d (0x%x)\n", p.Unk10, p.Unk10)
	s += fmt.Sprintf("\tUnknown uint32_t 11: %d (0x%x)\n", p.Unk11, p.Unk11)
	s += fmt.Sprintf("\tUnknown uint32_t 12: %d (0x%x)\n", p.Unk12, p.Unk12)
	s += fmt.Sprintf("\tUnknown uint32_t 13: %d (0x%x)\n", p.Unk13, p.Unk13)
	s += fmt.Sprintf("\tUnknown uint32_t 14: %d (0x%x)\n", p.Unk14, p.Unk14)
	s += fmt.Sprintf("\tUnknown uint32_t 15: %d (0x%x)\n", p.Unk15, p.Unk15)
	s += fmt.Sprintf("\tUnknown uint32_t 16: %d (0x%x)\n", p.Unk16, p.Unk16)
	s += fmt.Sprintf("\tUnknown uint32_t 17: %d (0x%x)\n", p.Unk17, p.Unk17)
	s += fmt.Sprintf("\tUnknown uint32_t 18: %d (0x%x)\n", p.Unk18, p.Unk18)
	s += fmt.Sprintf("\tUnknown String: %s\n", hex.EncodeToString(p.Unk19[:]))

	return s
}

type ScoreBoardPacket struct {
	PacketHead  uint32      //    0x00 0x0000001
	PacketType  uint32      //    0x04 0x100000E
	SongID      uint32      //    0x08
	ChartLevel  uint16      //    0x0C
	Type        uint8       //    0x0E
	Flag        uint8       //    0x0F
	Score       uint32      //    0x10
	RealScore0  uint32      //    0x14
	Unk0        [16]uint8   //    0x18
	RealScore1  uint32      //    Same as SongScore0, dafuq?
	Grade       uint32      //    0x2C
	Kcal        float32     //    0x30
	Perfect     uint32      //    0x34
	Great       uint32      //    0x38
	Good        uint32      //    0x3c
	Bad         uint32      //    0x40
	Miss        uint32      //    0x44
	MaxCombo    uint32      //    0x48
	EXP         uint16      //    0x4c
	PP          uint16      //    0x4e
	RunningStep uint16      //    0x50
	Unk2        uint16      //    0x52
	Unk3        uint32      //    0x54
	Unk4        uint32      //    0x58
	Unk5        uint32      //    0x5c
	RushSpeed   float32     //    0x60
	GameVersion PIUString12 //    0x64
	MachineID   uint32      //    0xFFFFFF
	ProfileID   uint32      //   0xB21
}

type LoginPacket struct {
	PacketHead  uint32      //    0x00 0x0000001
	PacketType  uint32      //    0x04 0x1000003
	PlayerID    uint32      //    0x08
	MachineID   uint32      //    0x0C
	AccessCode  PIUString32 //  Hex String
	PacketTrail uint32
}

type LoginPacketV2 struct {
	PacketHead  uint32      //    0x00 0x0000004
	PacketType  uint32      //    0x04 0x1000015
	PlayerID    uint32      //    0x08
	MachineID   uint32      //    0x0C
	AccessCode  PIUString32 //  Hex String
	Unk0        uint32
	GameVersion PIUString12
	PacketTrail uint32
}

type ProfilePacket struct {
	PacketHead  uint32      //    0x00 0x0000001
	PacketType  uint32      //    0x04 0x1000004
	PlayerID    uint32      //    0x08
	AccessCode  PIUString32 //    0x0C
	Unk0        uint32
	Nickname    PIUNickname //    0x30
	ProfileID   uint32      //    0x10
	CountryID   uint8       //    0x3C
	Avatar      uint8       //    0x40
	Level       uint8       //    0x42
	Unk1        uint8
	EXP         uint64
	PP          uint64
	RankSingle  uint64
	RankDouble  uint64
	RunningStep uint64
	PlayCount   uint32
	Kcal        float32
	Modifiers   uint64
	Unk2        uint32
	RushSpeed   float32
	Unk3        uint32
	Scores      [4384]UScore //    0x88
}

type UScore struct {
	SongID       uint32 //    0x00
	ChartLevel   uint8  //    0x04
	Unk0         uint8  //    0x05
	GameDataFlag uint16 //    0x06
	Score        uint32 //    0x08
	RealScore    uint32 //   Maybe
	Unk2         uint32 //    0x10
}

type RequestLevelUpInfoPacket struct {
	PacketHead uint32 //    0x00 0x0000001
	PacketType uint32 //    0x04 0x100000C
	ProfileID  uint32 //    0x08 0xBD5
}

type LevelUpInfoPacket struct {
	PacketHead uint32 //    0x00 0x0000001
	PacketType uint32 //    0x04 0x100000D
	ProfileID  uint32 //    0x08 0xBD5
	Level      uint32 //    0x0C
}

type GameOverPacket struct {
	PacketHead uint32 //    0x00 0x0000001
	PacketType uint32 //    0x04 0x1000001
	Unk0       uint32 //    0x08 0xBD5
}

type WorldBestPacket struct {
	PacketHead  uint32 //    0x00 0x00000002
	PacketType  uint32 //    0x04 0x10000009
	Unk0        uint32 //    0x08 5056
	Unk1        uint32 //    0x0C 0x0000000F
	Unk2        uint32 //    0x10 674200
	Unk3        uint32 //    0x14 0x00000000
	Unk4        uint32 //    0x18 0x00000000
	WorldScores [4095]WorldBestScore
	Unk5        uint32 //    0x?? 0x00000000
	Unk6        uint32 //    0x?? 0x00000000
	PacketTrail uint32 //    0x?? 0x00000000
}

type WorldBestScore struct {
	SongID     uint32      //
	ChartLevel uint16      //
	ChartMode  uint16      //
	Score      uint32      //
	Unk0       uint32      //
	Unk1       uint32      //
	Nickname   PIUNickname //
}

type RankModePacket struct {
	SimplePacket
	Ranks [400]SongRank
}

type SongRank struct {
	SongID uint32
	First  PIUNickname
	Second PIUNickname
	Third  PIUNickname
}

var ACKPacketLength = int(binary.Size(ACKPacket{}))
var MachineInfoPacketLength = int(binary.Size(MachineInfoPacket{}))
var ScoreBoardPacketLength = int(binary.Size(ScoreBoardPacket{}))
var LoginPacketLength = int(binary.Size(LoginPacket{}))
var ProfilePacketLength = int(binary.Size(ProfilePacket{}))
var RequestLevelUpInfoPacketLength = int(binary.Size(RequestLevelUpInfoPacket{}))
var LevelUpInfoPacketLength = int(binary.Size(LevelUpInfoPacket{}))
var GameOverPacketLength = int(binary.Size(GameOverPacket{}))
var WorldBestPacketLength = int(binary.Size(WorldBestPacket{}))
var ProfileBusyPacketLength = int(binary.Size(ProfileBusyPacket{}))
var ByePacketLength = int(binary.Size(ByePacket{}))
var EnterProfilePacketLength = int(binary.Size(EnterProfilePacket{}))
var RequestWorldBestPacketLength = int(binary.Size(RequestWorldBestPacket{}))
var RankModePacketLength = int(binary.Size(RankModePacket{}))
var RequestRankModePacketLength = int(binary.Size(RequestRankModePacket{}))
var LoginPacketV2Length = int(binary.Size(LoginPacketV2{}))
var KeepAlivePacketLength = int(binary.Size(KeepAlivePacket{}))

var BiggestPacket = 0

func init() {

	fmt.Println(ProfilePacketLength)

	packetLens := []int{
		ACKPacketLength, MachineInfoPacketLength, MachineInfoPacketLength,
		ScoreBoardPacketLength, LoginPacketLength, ProfilePacketLength, RequestLevelUpInfoPacketLength,
		LevelUpInfoPacketLength, GameOverPacketLength, WorldBestPacketLength,
		ProfileBusyPacketLength, ByePacketLength, EnterProfilePacketLength,
		RequestWorldBestPacketLength, RankModePacketLength, RequestRankModePacketLength,
		LoginPacketV2Length, KeepAlivePacketLength,
	}

	for _, v := range packetLens {
		if v > BiggestPacket {
			BiggestPacket = v
		}
	}
}
