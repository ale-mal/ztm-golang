//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printStatuses(statuses map[string]int) {
	fmt.Println("Total num is", len(statuses))

	stats := map[int]int{}
	for _, status := range statuses {
		stats[status]++
	}

	fmt.Println("Online num is", stats[Online])
	fmt.Println("Offline num is", stats[Offline])
	fmt.Println("Maintenance num is", stats[Maintenance])
	fmt.Println("Retired num is", stats[Retired])

	fmt.Println()
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	serverStatus := make(map[string]int)
	for _, server := range servers {
		serverStatus[server] = Online
	}

	printStatuses(serverStatus)

	serverStatus["darkstar"] = Retired
	serverStatus["aiur"] = Offline

	printStatuses(serverStatus)

	for server, _ := range serverStatus {
		serverStatus[server] = Maintenance
	}

	printStatuses(serverStatus)
}
