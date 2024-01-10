package cmd

import (
	// "api/libs/fakturownia/clent"
	fakturownia_type "api/libs/fakturownia"
	"api/libs/fakturownia/warehouse/all"
	"api/libs/ic/csv"
	"api/libs/ic/import"
	"api/server"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)
var rootCmd = &cobra.Command{
    Use:   "nazwaAplikacji",
    Short: "Krótki opis Twojej aplikacji",
    Long:  `Długi opis Twojej aplikacji...`,
    // Tutaj można dodać logikę, która będzie wykonywana, gdy uruchomiono bez podkomend
}
// serveCmd reprezentuje komendę serve
var ServeCmd = &cobra.Command{
    Use:   "serve",
    Short: "Uruchamia serwer HTTP",
    Long: `Uruchamia serwer HTTP dla aplikacji.`,
    Run: func(cmd *cobra.Command, args []string) {
        server.Server()
    },
}

var IcFileCSV = &cobra.Command{
    Use:   "icCSV",
    Short: "Wgrywanie pliku CSV z IC",
    Long: `Wgrywanie pliku CSV z IC`,
    Run: func(cmd *cobra.Command, args []string) {
        importCSV.ImportCSV()
    },
}
var IcFileSplit = &cobra.Command{
    Use:   "icSplit",
    Short: "Dzielenie pliku CSV z IC",
    Long: `Dzielenie pliku CSV z IC`,
    Run: func(cmd *cobra.Command, args []string) {
        spliter.Spliter()
    },   
}

var AllDock = &cobra.Command{
    Use:  "allDock",
    Short: "Wszystkie dokumenty z Fakturowni",
    Long: `Wszystkie dokumenty z Fakturowni`,
    Run: func(cmd *cobra.Command, args []string) {
        // docsFakturownia.AllDock(2, 10)
        payload := fakturownia_type.Payload {
            APIToken: os.Getenv("FAKTUROWNIA_TOKEN_API"),
            WarehouseDocument: fakturownia_type.WarehouseDocument{
                Kind:           fakturownia_type.KindPZ,
                Number:         nil,
                WarehouseID:    os.Getenv("FAKTUROWNIA_WAREHOUSE_ID"),
                IssueDate:      "2017-10-23",
                DepartmentName: "Department1 SA",
                ClientName:     "INTER CARS SPÓŁKA AKCYJNA",
                WarehouseActions: []fakturownia_type.WarehouseAction{
                    {
                        ProductName:      "Produkt A1",
                        PurchaseTax:      23,
                        PurchasePriceNet: 10.23,
                        Quantity:         1,
                    },
                    {
                        ProductName:      "Produkt A2",
                        PurchaseTax:      0,
                        PurchasePriceNet: 50,
                        Quantity:         2,
                    },
                },
            },
        }

        req, err := facturownia.AddWarehouse(payload)
        if err != nil {
            log.Println("Wystąpił błąd:", err)
            return
        }
        log.Println("Odpowiedź JSON:", string(req))
    },
}

func init() {
    rootCmd.AddCommand(ServeCmd)
    rootCmd.AddCommand(IcFileCSV)
    rootCmd.AddCommand(IcFileSplit)
    rootCmd.AddCommand(AllDock)
}

// Execute uruchamia główną komendę
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}