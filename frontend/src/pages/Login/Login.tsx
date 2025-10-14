import { Button, TextInput } from '@mantine/core'
import { AtSignIcon, Eye, EyeClosed } from 'lucide-react'
import { useState } from 'react'
import { actions, form, header, paper, title, wrapper } from './styles.css'

export default function Login() {
    const [showPassword, setShowPassword] = useState(false)

    const togglePasswordVisibility = () => {
        setShowPassword(!showPassword)
    }

    return (
        <div className={wrapper}>
            <div className={paper}>
                <div className={header}>
                    <h1 className={title}>Login</h1>
                </div>
                <div className={form}>
                    <TextInput
                        variant="filled"
                        label="Email"
                        placeholder="Email"
                        radius={'md'}
                        type="email"
                        rightSectionPointerEvents={'all'}
                        rightSection={
                            <AtSignIcon size={18} strokeWidth={2.5} />
                        }
                        //error="Invalid email"
                        required
                    />
                    <TextInput
                        variant="filled"
                        label="Password"
                        placeholder="NiceSecure123$%&"
                        radius={'md'}
                        type={showPassword ? 'text' : 'password'}
                        //error="Invalid password"
                        rightSection={
                            <>
                                {showPassword ? (
                                    <Eye
                                        size={18}
                                        strokeWidth={2.5}
                                        onClick={togglePasswordVisibility}
                                        style={{ cursor: 'pointer' }}
                                    />
                                ) : (
                                    <EyeClosed
                                        size={18}
                                        strokeWidth={2.5}
                                        onClick={togglePasswordVisibility}
                                        style={{ cursor: 'pointer' }}
                                    />
                                )}
                            </>
                        }
                        required
                    />
                    <div className={actions}>
                        <Button variant="gradient" size="md" fullWidth>
                            Login
                        </Button>
                    </div>
                </div>
            </div>
        </div>
    )
}
