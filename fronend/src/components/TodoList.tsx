import { useQuery } from '@tanstack/react-query';
import { getTodos } from '../lib/todo-api';

import { TodoItem } from './TodoItem';
import { TodoSchema } from '../lib/types';


export const TodoList = () => {
  const { status, data } = useQuery<TodoSchema[]>(
    {
      queryKey: ['todos'],
      queryFn: getTodos,
      networkMode: 'online',
      retry: 10,
      staleTime: 10,
      refetchInterval: 60 * 1000,
    });

  if (status === 'loading') return <h1>Loading...</h1>;

  const renderTodoItems = (completed: boolean) => {
    let filtered = data
      ?.filter((todo: TodoSchema) => todo.completed === completed)
      .map((todo: TodoSchema) => <TodoItem todo={todo} id={todo.id} key={todo.id} />);
    if (completed) {
      return filtered?.slice(0, 10)
    }
    return filtered
  };

  return (
    <div className='grid grid-cols-2 gap-4 left-0'>
      <div className='flex flex-col'>
        <strong>Todos</strong>
        {renderTodoItems(false)}
      </div>
      <div className='flex flex-col'>        
        <strong>Done</strong>
        {renderTodoItems(true)}</div>
    </div>
  );
};

