CREATE TABLE IF NOT EXISTS session (
	"id" uuid NOT NULL PRIMARY KEY, 
	"refresh_token" text NOT NULL, 
	"provider_refresh_token" text NOT NULL, 
	"client_ip" text NOT NULL, 
	"client_agent" text NOT NULL, 
	"expires" timestamp(6) NOT NULL, 
	"created_at" timestamp(6) NOT NULL, 
	"created_by" uuid,

  FOREIGN KEY (created_by)
    REFERENCES account(id)
)