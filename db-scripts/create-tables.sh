
export PGPASSWORD=$1

# Create database 
psql -h localhost -d postgres -U XPTO -p 5432 -a -q -f db-scripts/store.sql

# Create basic tables
psql -h localhost -d album_store -U XPTO -p 5432 -a -q -f api/album/album.sql
psql -h localhost -d album_store -U XPTO -p 5432 -a -q -f api/client/client.sql
psql -h localhost -d album_store -U XPTO -p 5432 -a -q -f api/transaction/transaction.sql
