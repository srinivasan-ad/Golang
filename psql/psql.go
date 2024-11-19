package main
import(
    "context"
	"fmt"
	"os"
   "github.com/jackc/pgx/v5/pgxpool"
)	

func main(){
	const connectionString = "postgresql://postgres.ejspmxevlxzxowxttbmy:Vilasbhai94!@aws-0-ap-south-1.pooler.supabase.com:6543/postgres"
	 ctx := context.Background()

	pool,err := pgxpool.New(ctx,connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Connected to database :)")
	defer pool.Close()
}
