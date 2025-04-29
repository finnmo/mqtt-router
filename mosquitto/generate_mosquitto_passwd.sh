#!/bin/sh

# Set the default usernames and passwords (you can override these via environment variables)
CAMERA_418_USER=${CAMERA_418_USER:-"418_camera_user"}
CAMERA_418_PASS=${CAMERA_418_PASS:-"418_camera_password"}
CAMERA_431_USER=${CAMERA_431_USER:-"431_camera_user"}
CAMERA_431_PASS=${CAMERA_431_PASS:-"431_camera_password"}
LOCAL_USER=${LOCAL_USER:-"local_user"}
LOCAL_PASS=${LOCAL_PASS:-"local_password"}
TEST_USER=${TEST_USER:-"test_user"}
TEST_PASS=${TEST_PASS:-"test_password"}

# Create the Mosquitto password file
echo "Generating mosquitto password file with 4 users..."

# Ensure the password file directory exists
mkdir -p ./mosquitto/config

# Generate the password file for each user
mosquitto_passwd -b ./mosquitto/config/passwd "$CAMERA_418_USER" "$CAMERA_418_PASS"
mosquitto_passwd -b ./mosquitto/config/passwd "$CAMERA_431_USER" "$CAMERA_431_PASS"
mosquitto_passwd -b ./mosquitto/config/passwd "$LOCAL_USER" "$LOCAL_PASS"
mosquitto_passwd -b ./mosquitto/config/passwd "$TEST_USER" "$TEST_PASS"

echo "Password file generated at ./mosquitto/config/passwd"
