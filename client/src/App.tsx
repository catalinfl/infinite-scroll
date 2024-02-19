import { useEffect, useState } from 'react'
import './index.css'
import axios from 'axios'
import { QueryClient, QueryClientProvider, useQuery } from '@tanstack/react-query'
import FilePagination from './component/FilePagination'
import FileInfiniteQuery from './component/FileInfiniteQuery'

function App() {

  const queryClient = new QueryClient()

  return (
    <QueryClientProvider client={queryClient}>
      <FileInfiniteQuery />
    </QueryClientProvider>
  )
}



export default App
