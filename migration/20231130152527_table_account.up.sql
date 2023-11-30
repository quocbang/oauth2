CREATE TABLE account (
	"id" uuid NOT NULL PRIMARY KEY, 
	"name" text NOT NULL, 
	"email" text NOT NULL, 
	"provider" smallint NOT NULL, 
	"provider_user_id" text NOT NULL, 
	"image" text NOT NULL, 
	"created_at" timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
)