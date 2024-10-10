### After cloning this project, you can run this program by simply running this command: go run main.go

### If you wanna add points for a payer, you can send POST request to localhost:8000/add with params in this format: 
{
         "payer" : "DANNON",
         "points" : 5000,
         "timestamp" : "2020-11-02T14:00:00Z"
}

### If you wanna spend points, you can send POST request to localhost:8000/spend with params in this format:
{"points" : 5000}

### If you wanna get points balance, you can send GET request to localhost:8000/balance with no params.