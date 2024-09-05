# api-go


#### PGX - package info and helpful insights 

    https://donchev.is/post/working-with-postgresql-in-go-using-pgx/

# Create a local module and import it from other package/module to go.mod file 

    # Creates module
    go mod init api/cpg 
    
    # Installs required packages
    go mod tidy
    
    # Go to other package that is going to use the module"api/cpg"
    cd ...

    # Add reference to the module and redirect it to the folder desired/where the module is
    go mod edit -replace=api/album=../album

    # Request the creation of the module link/visability at the current project
    go get api/album
    go mod tidy