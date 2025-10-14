import { globalStyle } from '@vanilla-extract/css'
import { BackgroundGray } from './colors'

globalStyle('*', {
    boxSizing: 'border-box',
})

globalStyle(':root', {
    fontFamily: 'system-ui, Avenir, Helvetica, Arial, sans-serif',
    lineHeight: 1.5,
    fontWeight: 400,
    WebkitFontSmoothing: 'antialiased',
    MozOsxFontSmoothing: 'grayscale',
    backgroundColor: BackgroundGray,
})

globalStyle('a', {
    fontWeight: 500,
    color: '#646cff',
    textDecoration: 'inherit',
})

globalStyle('a:hover', {
    color: '#535bf2',
})

globalStyle('h1', {
    fontSize: '3.2em',
    lineHeight: 1.1,
})

globalStyle('button', {
    borderRadius: '8px',
    border: '1px solid transparent',
    padding: '0.6em 1.2em',
    fontSize: '1em',
    fontWeight: 500,
    fontFamily: 'inherit',
    backgroundColor: '#1a1a1a',
    cursor: 'pointer',
    transition: 'border-color 0.25s',
})

globalStyle('button:hover', {
    borderColor: '#646cff',
})

globalStyle('button:focus, button:focus-visible', {
    outline: '4px auto -webkit-focus-ring-color',
})

globalStyle('[data-mantine-color-scheme="light"] a:hover', {
    color: '#747bff',
})

globalStyle('[data-mantine-color-scheme="light"] button', {
    backgroundColor: '#f9f9f9',
})

export {}
