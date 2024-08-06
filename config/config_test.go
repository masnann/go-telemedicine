package config

import (
	_ "github.com/lib/pq"
)

// func Test_OpenConnection(t *testing.T) {

// 	t.Run("Success Case - Success Open Connection", func(t *testing.T) {
// 		err := OpenConnection()
// 		assert.Nil(t, err)
// 		assert.NotNil(t, DBConnection())
// 		CloseConnectionDB()
// 	})

// 	t.Run("Failure Case - Error Call Setup Connection", func(t *testing.T) {
// 		DBDriver = "invalid_driver"

// 		err := OpenConnection()
// 		assert.NotNil(t, err)
// 		assert.EqualError(t, err, "failed to create the database connection")
// 		assert.Nil(t, DBConnection())
// 	})

// 	t.Run("Failure Case - Error Call Ping", func(t *testing.T) {

// 		DBDriver = "postgres"
// 		DBUser = "validuser"
// 		DBPass = "validpass"
// 		DBName = "validdb"
// 		DBHost = "invalidhost"
// 		DBPort = "5432"
// 		SSLMode = "disable"

// 		err := OpenConnection()
// 		assert.NotNil(t, err)
// 		assert.Contains(t, err.Error(), "dial tcp")
// 	})
// }

// func TestGetEnv(t *testing.T) {
// 	os.Setenv("DB_DRIVER", "postgres")
// 	os.Setenv("DB_NAME", "testdb")

// 	result := GetEnv("DB_DRIVER")
// 	assert.Equal(t, "postgres", result, "Expected DB_DRIVER to be 'postgres'")

// 	result = GetEnv("NON_EXISTENT_VAR", "default_value")
// 	assert.Equal(t, "default_value", result, "Expected default_value when environment variable is not set")

// 	result = GetEnv("ANOTHER_NON_EXISTENT_VAR")
// 	assert.Equal(t, "", result, "Expected empty string when environment variable is not set and no default value is provided")

// }
