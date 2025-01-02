"use client"

import { useRouter } from 'next/navigation'

export default function Home() {
  const router = useRouter()
  
  const handleGoogleLogin = () => {
    window.location.href = 'http://localhost:8000/auth/google/login'
  }

  return (
    <div className="flex min-h-screen flex-col items-center justify-center">
      <button 
        onClick={handleGoogleLogin}
        className="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
      >
        Login with Google
      </button>
    </div>
  )
}
