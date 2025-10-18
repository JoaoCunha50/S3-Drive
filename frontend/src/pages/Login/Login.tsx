import { Button, Image, TextInput } from '@mantine/core'
import { AxiosError } from 'axios'
import { AtSignIcon, Eye, EyeClosed } from 'lucide-react'
import { useState } from 'react'
import { api } from '../../config/api'
import {
    actions,
    error,
    form,
    header,
    paper,
    title,
    wrapper,
} from './styles.css'

type FormData = {
    email: string
    username: string
    password: string
}

export default function Login() {
    const [showPassword, setShowPassword] = useState(false)
    const [errorMessage, setError] = useState<string | null>(null)
    const [formData, setFormData] = useState<FormData>({
        email: '',
        username: '',
        password: '',
    })

    async function onSubmit() {
        try {
            console.log(api)
            const response = await api.post('/users/login', { ...formData })
            if (response.status === 200) {
                console.log(response.data)
                window.location.href = '/'
            } else {
                setError('Invalid credentials')
            }
        } catch (error) {
            console.log(error)

            if (error instanceof AxiosError && error.status === 404) {
                console.log(error)
                setError('')
                setError('Invalid credentials')
            } else {
                setError('Something wrong happened...')
            }
        }
    }

    const togglePasswordVisibility = () => {
        setShowPassword(!showPassword)
    }

    return (
        <div className={wrapper}>
            <div className={paper}>
                <div className={header}>
                    <Image src="/S3Drive.png" h={100} w={100} />
                    <h1 className={title}>S3Drive</h1>
                </div>
                <div className={form}>
                    <TextInput
                        variant="filled"
                        label="Email or Username"
                        placeholder="johndoe(@gmail.com)"
                        radius={'md'}
                        type="email"
                        rightSectionPointerEvents={'all'}
                        rightSection={
                            <AtSignIcon size={18} strokeWidth={2.5} />
                        }
                        onChange={(e) =>
                            setFormData({
                                ...formData,
                                email: e.target.value,
                                username: e.target.value,
                            })
                        }
                        required
                    />
                    <TextInput
                        variant="filled"
                        label="Password"
                        placeholder="NiceSecure123$%&"
                        radius={'md'}
                        type={showPassword ? 'text' : 'password'}
                        onChange={(e) =>
                            setFormData({
                                ...formData,
                                password: e.target.value,
                            })
                        }
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
                    <h1 className={error}>{errorMessage}</h1>
                    <div className={actions}>
                        <Button
                            type="submit"
                            variant="gradient"
                            size="md"
                            fullWidth
                            onClick={onSubmit}
                        >
                            Login
                        </Button>
                    </div>
                </div>
            </div>
        </div>
    )
}
