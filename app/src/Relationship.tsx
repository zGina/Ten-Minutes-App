import React from 'react';
import { List, Datagrid,SelectInput , TextField,ReferenceField, EmailField,BooleanInput, Edit,SimpleForm,TextInput,DateInput,ArrayInput,SimpleFormIterator,Create } from 'react-admin';
import MyUrlField from './MyUrlField';

export const RelationshipList = (props: any) => {
  return (<List {...props}>
    <Datagrid rowClick="edit" >
      <TextField source="type" />
      <ReferenceField label="AttackPattern" source="source_ref" reference="attackPatterns">
      <TextField source="id" />
      </ReferenceField>

      <ReferenceField label="AttackPattern" source="target_ref" reference="attackPatterns">
      <TextField source="id" />
      </ReferenceField>
    </Datagrid>
  </List>);
};


const RelationshipForm =(props:any)=>{

  return (
 <SimpleForm warnWhenUnsavedChanges {...props}>
         {/* <TextInput source="id" /> */}
          <TextInput source="id_" />
          <TextInput source="type" />

          <ReferenceField label="AttackPattern" source="source_ref" reference="attackPatterns">
          <SelectInput source="name" />
      </ReferenceField>

      <ReferenceField label="AttackPattern" source="target_ref" reference="attackPatterns">
          <SelectInput source="name" />
      </ReferenceField>

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