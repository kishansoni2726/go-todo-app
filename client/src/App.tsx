import { Button, Container, Stack } from '@chakra-ui/react'
import './App.css'
import NavBar from './components/NavBar';
import TodoForm from './components/TodoForm'
import TodoList from './components/TodoList'
import React from 'react';

export const BASE_URL = "http://localhost:5000/api" ;
function App() {

  return(
    <Stack h="100vh">
      <NavBar></NavBar>

      <Container>

        <TodoForm>
        </TodoForm>

        <TodoList>
        </TodoList>
        
      </Container>
    </Stack>
  );
  
}

export default App;