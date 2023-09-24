import './index.css'
import { TodoList } from './components/TodoList'
import { TodoForm } from './components/TodoForm'
import {
  QueryClient,
  QueryClientProvider,
} from '@tanstack/react-query'
import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
import { ToastContainer } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 5,
      refetchOnWindowFocus: false,
      refetchOnMount: false,
      retry: 1,
    },
  },
})
function App() {

  return (
    <main className="flex h-screen justify-center">
      <div className="w-full p-6 m-10 bg-[#434654] rounded-md shadow-md lg:max-w-xl">
        <div className="flex flex-col overflow-auto">

          <QueryClientProvider client={queryClient}>
            <TodoForm />
            <TodoList />
            <ReactQueryDevtools initialIsOpen={true} />
          </QueryClientProvider>
          <ToastContainer />
        </div>
      </div>
    </main>
  )
}

export default App
