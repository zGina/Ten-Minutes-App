import React from 'react';
import { List, Datagrid, UrlField, TextField, EmailField,BooleanInput, Edit,SimpleForm,TextInput,DateInput,ArrayInput,SimpleFormIterator,Create } from 'react-admin';
import MyUrlField from './MyUrlField';

export const AttackPatternList = (props: any) => {
  return (<List {...props}>
    <Datagrid >
      <TextField source="name" />
      <TextField label="STIX_ID" source="external_references[0].external_id" />
      <UrlField source="external_references[0].url" />
    </Datagrid>
  </List>);
};

const AttackPatternForm =(props:any)=>{

  return (
 <SimpleForm warnWhenUnsavedChanges {...props}>
          <TextInput source="name" />
          <TextInput source="description" />
          <TextInput source="id" />
          <TextInput source="type" />

        <ArrayInput source='kill_chain_phases'>
        <SimpleFormIterator>
        <TextInput source="kill_chain_name"  label="kill_chain_name"/>
        <TextInput source="phase_name"  label="phase_name"/>
        </SimpleFormIterator>
        </ArrayInput>

        <ArrayInput source='external_references'>
        <SimpleFormIterator>
        <TextInput source="source_name"  label="source_name"/>
        <TextInput source="external_id"  label="external_id"/>
        <TextInput source="url"  label="url"/>
        </SimpleFormIterator>
        </ArrayInput>

        <TextInput source="x_mitre_version" label="x_mitre_version" title="x_mitre_version"  />
        <TextInput source="x_mitre_detection" label="x_mitre_detection" title="x_mitre_detection"  />

        <ArrayInput label  = "x_mitre_platforms" source="x_mitre_platforms" >
  <SimpleFormIterator >
        <TextInput  />
  </SimpleFormIterator >
        </ArrayInput>

        <ArrayInput label  = "x_mitre_permissions_required" source="x_mitre_permissions_required" >
  <SimpleFormIterator >
        <TextInput  />
  </SimpleFormIterator >
        </ArrayInput>

        <ArrayInput label  = "x_mitre_data_sources" source="x_mitre_data_sources" >
  <SimpleFormIterator >
        <TextInput  />
  </SimpleFormIterator >
        </ArrayInput>

    <BooleanInput source="x_mitre_is_subtechnique" label="x_mitre_is_subtechnique" />
</SimpleForm>
  )
}
export const AttackPatternEdit = (props:any) => (
 <Edit title="编辑AttackPattern"  {...props} >
      {<AttackPatternForm/>}
</Edit>
);

export const AttackPatternCreate = (props:any) => (
   <Create title="新建一个AttackPattern" {...props}>
         {<AttackPatternForm/>}
   </Create>
);