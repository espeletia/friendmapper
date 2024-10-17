import React from "react";

// Define styles for the profile picture container and the plus icon
const styles = {
  container: {
    display: 'flex',
    flexDirection: 'column',
    alignItems: 'center',
    gap: '10px', // Space between label and circle
  } as React.CSSProperties,

  circle: {
    width: '80px', // Size of the circle
    height: '80px', // Same height as width to make it a circle
    borderRadius: '50%', // Full circle
    backgroundColor: '#f0f0ff', // Light purple background
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    cursor: 'pointer', // Pointer cursor on hover
  } as React.CSSProperties,

  plusIcon: {
    fontSize: '24px', // Size of the plus icon
    color: '#6C63FF', // Purple color for the plus
  } as React.CSSProperties,

  label: {
    fontSize: '14px',
    fontWeight: 'bold',
    color: '#333333', // Dark text for the label
  } as React.CSSProperties
};

interface ProfilePicturePlaceholderProps {
  label?: string;
  circleColor?: string;
  iconColor?: string;
}

const ProfilePicturePlaceholder = (props: ProfilePicturePlaceholderProps) => {
  const {
    label = 'Profilová fotka', // Default label text
    circleColor = '#f0f0ff', // Default background color for the circle
    iconColor = '#6C63FF', // Default color for the plus icon
  } = props;

  return (
    <div style={styles.container}>
      <span style={styles.label}>{label}</span>
      <div
        style={{
          ...styles.circle,
          backgroundColor: circleColor, // Dynamic background color for the circle
        }}
      >
        <span
          className="material-icons"
          style={{
            ...styles.plusIcon,
            color: iconColor, // Dynamic color for the plus icon
          }}
        >
          add
        </span>
      </div>
    </div>
  );
};

export default ProfilePicturePlaceholder;
