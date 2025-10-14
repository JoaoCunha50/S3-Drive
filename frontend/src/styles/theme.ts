import { createTheme, type CSSVariablesResolver } from '@mantine/core'
import { themeToVars } from '@mantine/vanilla-extract'
import * as colors from './colors'

export const theme = createTheme({
    fontFamily: 'system-ui, Avenir, Helvetica, Arial, sans-serif',
    white: colors.White,
    cursorType: 'pointer',
    primaryColor: 'primary',
    colors: {
        primary: [
            colors.Indigo,
            colors.Indigo,
            colors.Indigo,
            colors.Indigo,
            colors.Indigo,
            colors.Indigo,
            colors.Indigo,
            colors.Indigo,
            colors.Indigo,
            colors.Indigo,
        ],
    },
    focusRing: 'auto',
    defaultRadius: 'lg',
    defaultGradient: {
        from: colors.Indigo,
        to: colors.Violet,
        deg: 45,
    },
    breakpoints: {
        mobile: '600px',
        tablet: '768px',
        desktop: '1024px',
    },
})

export const variableResolver: CSSVariablesResolver = () => ({
    variables: {
        '--mantine-color-body': colors.BackgroundGray,
    },
    light: {
        '--mantine-color-body': colors.White,
    },
    dark: {
        '--mantine-color-body': colors.BackgroundGray,
    },
})

export const vars = themeToVars(theme)
