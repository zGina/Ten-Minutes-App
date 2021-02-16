import React from 'react';
import { List, Datagrid, TextField, EmailField,Edit,SimpleForm,TextInput,DateInput,ArrayInput,SimpleFormIterator,Create } from 'react-admin';
import MyUrlField from './MyUrlField';

export const UserList = (props: any) => {
  return (<List {...props}>
    <Datagrid rowClick="edit">
      <TextField source="id" />
      <TextField source="name" />
      <EmailField source="email" />
      <TextField source="phone" />
      <MyUrlField source="website" />
      <TextField source="company.name" />
    </Datagrid>
  </List>);
};


const UserForm =(props:any)=>{

  return (
 <SimpleForm warnWhenUnsavedChanges {...props}>
         {/* <TextInput source="id" /> */}
          <TextInput source="name" />
          <TextInput source="username" />
          <TextInput source="email" />

    <TextInput source="address.street" label="street" placeholder="street" title="target_id"/>
    <TextInput source="address.suite" label="suite" placeholder="suite" title="uri" />
    <TextInput source="address.city" label="city" title="title"  />
    <TextInput source="address.zipcode" label="zipcode" title="title"  />
    <TextInput source="address.geo.lat" label="lat" title="lat"  />
    <TextInput source="address.geo.lng" label="lng" title="lng"  />
          <TextInput source="phone" />
          <TextInput source="website" />
    <TextInput source="company.name" label="name"  />
    <TextInput source="company.catchPhrase" label="catchPhrase" />
    <TextInput source="company.bs" label="BS"  />
</SimpleForm>
  )
}
export const UserEdit = (props:any) => (
 <Edit title="编辑User"  {...props} >
      {<UserForm/>}
</Edit>
);

export const UserCreate = (props:any) => (
   <Create title="新建一个User" {...props}>
         {<UserForm/>}
   </Create>
);