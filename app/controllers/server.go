package controllers

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/raggaer/tiger/app/models"
)

type status struct {
	Online        bool
	Address       string
	Name          string
	LoginPort     string
	OwnerEmail    string
	OwnerName     string
	Motd          string
	Location      string
	URL           string
	Uptime        uint64
	PlayersPeak   uint32
	MapName       string
	MapAuthor     string
	MapWidth      uint16
	MapHeight     uint16
	PlayersMax    uint32
	PlayersOnline uint32
	UptimeHours   uint64
	UptimeMinutes uint64
}

const (
	basicServerInfo     = 1
	ownerServerInfo     = 2
	miscServerInfo      = 4
	playersInfo         = 8
	mapInfo             = 16
	extendedPlayersInfo = 32
	playerStatusInfo    = 64
)

func retrieveServerStatus(address string, info uint16) ([]byte, error) {
	b := []byte{}
	buff := bytes.NewBuffer(b)

	// Write protocol identifier
	if err := binary.Write(buff, binary.LittleEndian, byte(0xFF)); err != nil {
		return nil, err
	}

	// Write protocol byte switch
	if err := binary.Write(buff, binary.LittleEndian, byte(0x01)); err != nil {
		return nil, err
	}

	// Write requested information
	if err := binary.Write(buff, binary.LittleEndian, info); err != nil {
		return nil, err
	}

	finalBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(finalBuffer, binary.LittleEndian, uint16(buff.Len())); err != nil {
		return nil, err
	}
	if err := binary.Write(finalBuffer, binary.LittleEndian, []byte(buff.Bytes())); err != nil {
		return nil, err
	}

	// Create connection to the server
	conn, err := net.DialTimeout("tcp", address, time.Second*2)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// Send message to the server
	if _, err := conn.Write(finalBuffer.Bytes()); err != nil {
		return nil, err
	}

	// Create reading buffer
	readBuffer := make([]byte, 2024)
	bytesNumber, err := conn.Read(readBuffer)
	if err != nil {
		return nil, err
	}
	return readBuffer[:bytesNumber], nil
}

func parseServerStatus(s []byte) (*status, error) {
	buff := bytes.NewBuffer(s)
	info := &status{}

	for {
		//Read header
		header, err := buff.ReadByte()
		if err == io.EOF {
			break
		}

		switch header {
		// Case BASIC_SERVER_INFO
		case 0x10:

			// Read server name
			serverName, err := readString(buff)

			if err != nil {
				return nil, errors.New("Cannot read server name")
			}

			// Read server address

			serverAddress, err := readString(buff)

			if err != nil {
				return nil, errors.New("Cannot read server address")
			}

			// Read server login port
			loginPort, err := readString(buff)

			if err != nil {
				return nil, errors.New("Cannot read server login port")
			}

			info.Address = serverAddress
			info.Name = serverName
			info.LoginPort = loginPort

		// Case OWNER_INFO
		case 0x11:

			// Read owner name
			ownerName, err := readString(buff)

			if err != nil {
				return nil, errors.New("Cannot read server owner name")
			}

			// Read owner email
			ownerEmail, err := readString(buff)

			if err != nil {
				return nil, errors.New("Cannot read server owner email")
			}

			info.OwnerEmail = ownerEmail
			info.OwnerName = ownerName

		// Case MISC_INFO
		case 0x12:

			// Get MOTD
			serverMotd, err := readString(buff)

			if err != nil {
				return nil, errors.New("Cannot read server message of the day")
			}

			// Get server location
			serverLocation, err := readString(buff)

			if err != nil {
				return nil, errors.New("Cannot read server location")
			}

			// Get server URL
			serverURL, err := readString(buff)

			if err != nil {
				return nil, errors.New("Cannot read server URL")
			}

			// Get server uptime
			var serverUptime uint64

			if err := binary.Read(buff, binary.LittleEndian, &serverUptime); err != nil {
				return nil, errors.New("Cannot read server uptime")
			}

			info.Motd = serverMotd
			info.Location = serverLocation
			info.URL = serverURL
			info.Uptime = serverUptime

			// Case PLAYERS_INFO
		case 0x20:

			// Get players online
			var playersOnline uint32

			if err := binary.Read(buff, binary.LittleEndian, &playersOnline); err != nil {
				return nil, errors.New("Cannot read server players online")
			}

			// Get max players
			var playersMax uint32

			if err := binary.Read(buff, binary.LittleEndian, &playersMax); err != nil {
				return nil, errors.New("Cannot read server maximum players")
			}

			// Get players peak
			var playersPeak uint32

			if err := binary.Read(buff, binary.LittleEndian, &playersPeak); err != nil {
				return nil, errors.New("Cannot read server players peak")
			}

			info.PlayersOnline = playersOnline
			info.PlayersMax = playersMax
			info.PlayersPeak = playersPeak

			// Case MAP_INFO
		case 0x30:

			// Get map name
			mapName, err := readString(buff)

			if err != nil {
				return nil, errors.New("Cannot read server map name")
			}

			// Get map author
			mapAuthor, err := readString(buff)

			if err != nil {
				return nil, errors.New("Cannot read server map author")
			}

			// Get map width
			var mapWidth uint16

			if err := binary.Read(buff, binary.LittleEndian, &mapWidth); err != nil {
				return nil, errors.New("Cannot read server map width")
			}

			// Get map height
			var mapHeight uint16

			if err := binary.Read(buff, binary.LittleEndian, &mapHeight); err != nil {
				return nil, errors.New("Cannot read server map height")
			}

			info.MapName = mapName
			info.MapAuthor = mapAuthor
			info.MapWidth = mapWidth
			info.MapHeight = mapHeight
		}
	}
	return info, nil
}

func readString(buff *bytes.Buffer) (string, error) {
	var length uint16

	if err := binary.Read(buff, binary.LittleEndian, &length); err != nil {
		return "", err
	}

	b := make([]byte, length)

	if err := binary.Read(buff, binary.LittleEndian, &b); err != nil {
		return "", err
	}
	return string(b), nil
}

// ServerStatus retrieves the current server status
func ServerStatus(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Retrieve server status
	var currentServerStatus *status
	svStatus, found := context.Cache.Get("serverStatus")
	if found {
		s, ok := svStatus.(*status)
		if !ok {
			return nil, errors.New("Unable to convert server status interface to status pointer")
		}
		currentServerStatus = s
	} else {
		// Retrieve server information
		b, err := retrieveServerStatus(
			context.Config.Server.Address,
			basicServerInfo|ownerServerInfo|miscServerInfo|playersInfo|mapInfo,
		)

		// If there is an error
		// Render message as server offline
		if err != nil {
			data, err := context.ExecuteTemplate("server_status", map[string]interface{}{
				"status": &status{
					Online: false,
				},
			})
			if err != nil {
				return nil, err
			}

			return &discordgo.MessageEmbed{
				Title:       "Server status",
				Description: data,
				Color:       3447003,
			}, nil
		}

		// Retrieve server status
		s, err := parseServerStatus(b)
		if err != nil {
			return nil, err
		}
		currentServerStatus = s
		currentServerStatus.Online = true
		currentServerStatus.UptimeHours = currentServerStatus.Uptime / 3600
		currentServerStatus.UptimeMinutes = (currentServerStatus.Uptime - (3600 * currentServerStatus.UptimeHours)) / 60

		// Set server status cache
		context.Cache.Set("serverStatus", s, 5*time.Minute)
	}

	data, err := context.ExecuteTemplate("server_status", map[string]interface{}{
		"status": currentServerStatus,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Server status",
		Description: data,
		Color:       3447003,
	}, nil
}

// LatestDeaths retrieves the server latest deaths
func LatestDeaths(context *Context, s *discordgo.Session, m *discordgo.MessageCreate) (*discordgo.MessageEmbed, error) {
	// Load server latest deaths
	deaths, err := models.GetServerDeaths(context.DB, 10)
	if err != nil {
		return nil, err
	}

	data, err := context.ExecuteTemplate("server_death", map[string]interface{}{
		"deaths": deaths,
	})
	if err != nil {
		return nil, err
	}

	return &discordgo.MessageEmbed{
		Title:       "Latest deaths",
		Description: data,
		Color:       3447003,
	}, nil
}
