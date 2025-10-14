import { MantineProvider } from '@mantine/core'
import '@mantine/core/styles.css'
import AppRouter from './router/Router'
import { theme, variableResolver } from './styles/theme'

function App() {
    return (
        <MantineProvider
            theme={theme}
            cssVariablesResolver={variableResolver}
            defaultColorScheme="auto"
        >
            <AppRouter />
        </MantineProvider>
    )
}

export default App
