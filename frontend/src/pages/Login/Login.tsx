import { Button, Image, TextInput } from '@mantine/core'
import { AtSignIcon, Eye, EyeClosed } from 'lucide-react'
import { useContext, useState } from 'react'
import { useTranslation } from 'react-i18next'
import { useNavigate } from 'react-router'
import { UserContext } from '../../context/Context'
import { Login as LoginRequest } from '../../services/requests/Auth'
import { AuthTokenKey } from '../../services/storage/StorageKeys'
import {
    actions,
    error,
    form,
    header,
    paper,
    title,
    wrapper,
} from './LoginStyles.css'

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
    const navigate = useNavigate()
    const { t } = useTranslation()
    const { setUser } = useContext(UserContext)

    async function onSubmit() {
        try {
            const response = await LoginRequest(formData)
            if (response?.success) {
                localStorage.setItem(AuthTokenKey, response.data.token)
                setUser(response.data.user)
                navigate('/')
            } else {
                setError(response?.message!)
            }
        } catch (error) {
            console.log(error)
            setError('Something wrong happened...')
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
                    <h1 className={title}>R2Drive</h1>
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
                        placeholder="NiceSecure123$%"
                        radius={'md'}
                        type={showPassword ? 'text' : 'password'}
                        onChange={(e) =>
                            setFormData({
                                ...formData,
                                password: e.target.value,
                            })
                        }
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
                            {t('LOGIN')}
                        </Button>
                    </div>
                </div>
            </div>
        </div>
    )
}
