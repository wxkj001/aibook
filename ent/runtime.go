// Code generated by ent, DO NOT EDIT.

package ent

import (
	"aibook/ent/schema"
	"aibook/ent/user"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescRole is the schema descriptor for role field.
	userDescRole := userFields[5].Descriptor()
	// user.DefaultRole holds the default value on creation for the role field.
	user.DefaultRole = userDescRole.Default.(string)
	// userDescBan is the schema descriptor for ban field.
	userDescBan := userFields[6].Descriptor()
	// user.DefaultBan holds the default value on creation for the ban field.
	user.DefaultBan = userDescBan.Default.(bool)
	// userDescLoginTime is the schema descriptor for login_time field.
	userDescLoginTime := userFields[8].Descriptor()
	// user.DefaultLoginTime holds the default value on creation for the login_time field.
	user.DefaultLoginTime = userDescLoginTime.Default.(func() time.Time)
	// user.UpdateDefaultLoginTime holds the default value on update for the login_time field.
	user.UpdateDefaultLoginTime = userDescLoginTime.UpdateDefault.(func() time.Time)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[9].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[10].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
}
