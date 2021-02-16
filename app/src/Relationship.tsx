import React from 'react';
import {List, AutocompleteInput,Datagrid,SelectInput ,ReferenceInput, TextField,ReferenceField, EmailField,BooleanInput, Edit,SimpleForm,TextInput,DateInput,ArrayInput,SimpleFormIterator,Create } from 'react-admin';
import MyUrlField from './MyUrlField';
import { withStyles } from '@material-ui/core';

export const RelationshipList = (props: any) => {
  return (<List {...props}>
    <Datagrid rowClick="edit" >
      <TextField source="type" />

      <ReferenceField label="AttackPattern" source="source_ref" reference="attackPatterns"
       filterToQuery={ (searchText:any) => ({ name: searchText })}>
      <TextField source="name" />
      </ReferenceField>

      <ReferenceField label="AttackPattern" source="target_ref" reference="attackPatterns"
       filterToQuery={ (searchText:any) => ({ name: searchText })}>
      <TextField source="name" />
      </ReferenceField>
      
    </Datagrid>
  </List>);
};
// const AutocompleteArrayInputInDialog = withStyles({
//     suggestionsContainer: { zIndex: 2000 },
// })(AutocompleteArrayInput);

const RelationshipForm =(props:any)=>{

  return (
 <SimpleForm {...props}>
         {/* <TextInput source="id" /> */}
          <TextInput source="id" />
          <TextInput source="type" />

          <ReferenceInput label="source_ref" source="source_ref" reference="attackPatterns"   filterToQuery={ (searchText:any) => ({ name: searchText })}>
          <AutocompleteInput  source="name" optionText='name'  />

      </ReferenceInput>

      <ReferenceInput label="target_ref" source="target_ref" reference="attackPatterns">
          {/* <TextInput source="name" /> */}
          <AutocompleteInput  source="name" optionText='name'  />
      </ReferenceInput>

        <TextInput source="relationship_type"  label="relationship_type"/>
        <TextInput source="x_mitre_detection" label="x_mitre_detection" title="x_mitre_detection"  />

</SimpleForm>
  )
}
export const RelationshipEdit = (props:any) => (
 <Edit title="编辑Relationship"  {...props} >
      {<RelationshipForm/>}
</Edit>
);

export const RelationshipCreate = (props:any) => (
   <Create title="新建一个Relationship" {...props}>
         {<RelationshipForm/>}
   </Create>
);