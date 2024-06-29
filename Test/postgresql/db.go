package postgresql

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbURL = "postgres://postgres:password@localhost:5433/postgres?sslmode=disable"
)

type TDB struct {
	Pool *pgxpool.Pool
}

func NewTDB() *TDB {
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}
	return &TDB{Pool: pool}
}

func (d *TDB) SetUp(t *testing.T, tableName ...string) {
	t.Helper()
	d.truncateTable(context.Background(), tableName...)
}
func (d *TDB) Fill() {
	d.Pool.Exec(context.Background(), `
	INSERT INTO public.orders(orderid, addresseeid, shelflife, order_status, status_updated_date, hash_code, weight, price)
	VALUES 
	(1,1 ,'2024-01-30 16:00:00' ,0 ,'2024-06-30 16:00:00' ,'asdasdasdasd' ,50 ,101 );`)
	d.Pool.Exec(context.Background(), `
	INSERT INTO public.orders(orderid, addresseeid, shelflife, order_status, status_updated_date, hash_code, weight, price)
	VALUES 
	(2,2 ,'2024-11-30 16:00:00' ,0 ,'2024-06-30 16:00:00' ,'asdasdasdasd' ,50 ,101 );`)
	d.Pool.Exec(context.Background(), `
	INSERT INTO public.orders(orderid, addresseeid, shelflife, order_status, status_updated_date, hash_code, weight, price)
	VALUES 
	(3,2 ,'2024-11-30 16:00:00' ,1 ,'2024-06-30 16:00:00' ,'asdasdasdasd' ,50 ,101 );`)
	d.Pool.Exec(context.Background(), `
	INSERT INTO public.orders(orderid, addresseeid, shelflife, order_status, status_updated_date, hash_code, weight, price)
	VALUES 
	(4,1 ,'2024-06-30 16:00:00', 0 ,'2024-06-30 16:00:00' ,'asdasdasdasd' ,50 ,101 );`)
	d.Pool.Exec(context.Background(), `
	INSERT INTO public.orders(orderid, addresseeid, shelflife, order_status, status_updated_date, hash_code, weight, price)
	VALUES 
	(5,3 ,'2024-10-30 16:00:00' ,-1 ,'2024-06-30 16:00:00' ,'asdasdasdasd' ,50 ,101 );`)
	d.Pool.Exec(context.Background(), `
	INSERT INTO public.orders(orderid, addresseeid, shelflife, order_status, status_updated_date, hash_code, weight, price)
	VALUES 
	(6,3 ,'2024-10-30 16:00:00' ,-1 ,'2024-06-30 16:00:00' ,'asdasdasdasd' ,50 ,101 );`)
	d.Pool.Exec(context.Background(), `
	INSERT INTO public.orders(orderid, addresseeid, shelflife, order_status, status_updated_date, hash_code, weight, price)
	VALUES 
	(7,3 ,'2024-10-30 16:00:00' ,-1 ,'2024-06-30 16:00:00' ,'asdasdasdasd' ,50 ,101 );`)
	d.Pool.Exec(context.Background(), `
	INSERT INTO public.orders(orderid, addresseeid, shelflife, order_status, status_updated_date, hash_code, weight, price)
	VALUES 
	(8,3 ,'2024-10-30 16:00:00' ,-1 ,'2024-06-30 16:00:00' ,'asdasdasdasd' ,50 ,101 );`)
	d.Pool.Exec(context.Background(), `
	INSERT INTO public.orders(orderid, addresseeid, shelflife, order_status, status_updated_date, hash_code, weight, price)
	VALUES 
	(9,3 ,'2024-10-30 16:00:00' ,-1 ,'2024-06-30 16:00:00' ,'asdasdasdasd' ,50 ,101 );`)
	d.Pool.Exec(context.Background(), `
	INSERT INTO public.orders(orderid, addresseeid, shelflife, order_status, status_updated_date, hash_code, weight, price)
	VALUES 
	(10,3 ,'2024-10-30 16:00:00' ,-1 ,'2024-06-30 16:00:00' ,'asdasdasdasd' ,50 ,101 );`)
}

func (d *TDB) TearDown(t *testing.T) {
	t.Helper()

}

func (d *TDB) truncateTable(ctx context.Context, tableName ...string) {

	q := fmt.Sprintf("TRUNCATE %s", strings.Join(tableName, ","))
	if _, err := d.Pool.Exec(ctx, q); err != nil {
		panic(err)
	}
}
