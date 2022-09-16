import React from 'react'

interface SearchIconProps {
  width?: string
  height?: string
}

function SearchIcon(props: SearchIconProps) {
  let { width, height } = props
  width = width ? width : '24px'
  height = height ? height : '24px'
  return (
    <svg
      width={width}
      height={height}
      viewBox="0 0 24 24"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M9.22057 3C10.8938 3 12.4984 3.68482 13.6816 4.90381C14.8647 6.12279 15.5294 7.77609 15.5294 9.5C15.5294 11.11 14.9567 12.59 14.0153 13.73L14.2773 14H15.0441L19.897 19L18.4412 20.5L13.5882 15.5V14.71L13.3262 14.44C12.2197 15.41 10.7832 16 9.22057 16C7.54736 16 5.94269 15.3152 4.75955 14.0962C3.57642 12.8772 2.91174 11.2239 2.91174 9.5C2.91174 7.77609 3.57642 6.12279 4.75955 4.90381C5.94269 3.68482 7.54736 3 9.22057 3ZM9.22057 5C6.7941 5 4.85292 7 4.85292 9.5C4.85292 12 6.7941 14 9.22057 14C11.647 14 13.5882 12 13.5882 9.5C13.5882 7 11.647 5 9.22057 5Z"
        fill="black"
      />
    </svg>
  )
}

export default SearchIcon
