"use client"
import { useRouter } from 'next/navigation'
import { useEffect } from 'react'

export default function Callback() {
  const router = useRouter()

  useEffect(() => {
    const handleCallback = async () => {
      const response = await fetch('http://localhost:8000/auth/google/callback' + window.location.search)
      const data = await response.json()
      
      localStorage.setItem('token', data.token)
      router.push('/dashboard')
    }

    handleCallback()
  }, [])

  return (
    <div className="flex min-h-screen items-center justify-center">
      <div className="text-xl">Redirecting to dashboard...</div>
    </div>
  )
}


