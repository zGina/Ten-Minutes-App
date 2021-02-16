import React from 'react';
import PostIcon from '@material-ui/icons/Book';
import UserIcon from '@material-ui/icons/Group';
import { Admin, Resource,EditGuesser } from 'react-admin';
import jsonServerProvider from 'ra-data-json-server';
import Dashboard from './Dashboard';

import { PostList, PostEdit, PostCreate } from './Posts';
import { UserList,UserEdit,UserCreate } from './Users';

const dataProvider = jsonServerProvider("http://127.0.0.1:6868");
const Title = () => (<div>Mitre_attack</div>)

const App = () => (
  <Admin title={<Title/>} dashboard={Dashboard} dataProvider={dataProvider}>
      <Resource name="posts" list={PostList} edit={PostEdit} create={PostCreate} icon={PostIcon} />
      <Resource name="users" list={UserList} edit={UserEdit}  create={UserCreate} icon={UserIcon} />
  </Admin>
)

export default App;
