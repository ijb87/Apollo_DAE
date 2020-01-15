<template>
  <div id="app">
    <sorted-table :values="values" align="center">
      <thead>
        <tr>
          <th scope="col" style="width: 5rem;">
            <sort-link name="id">ID</sort-link>
          </th>
          <th scope="col" style="width: 10rem;">
            <sort-link name="name">Name</sort-link>
          </th>
          <th scope="col" style="width: 20rem;">
            <sort-link name="occupation">Occupation</sort-link>
          </th>
        </tr>
      </thead>
      <template #body="sort">
        <tbody>
          <tr v-for="value in sort.values" :key="value.id">
            <td>{{ value.id }}</td>
            <td>{{ value.name }}</td>
            <td>{{ value.occupation }}</td>
          </tr>
        </tbody>
      </template>
    </sorted-table>
  </div>
</template>

<script>
export default {
  name: "App",
  data: function() {
    return {
      values: [
        { name: "Aisling Barajas", id: 2, occupation: 'Public Health Training Assistant' },
        { name: "Dawid Wagner", id: 1, occupation: 'Acoustic Sensor Operator' },
        { name: "Menaal Mcdougall", id: 3, occupation: 'Transportation Manager' }
      ]
    }
  }
}
</script>

<style>
table {  
    color: #333; /* Lighten up font color */
    font-family: Helvetica, Arial, sans-serif; /* Nicer font */
    width: 640px; 
    border-collapse: 
    collapse; border-spacing: 0; 
}

td, th { 
  border: 1px solid #CCC; 
  height: 30px; 
  }

th {  
    background: #F3F3F3;
    font-weight: bold;
    text-align: center;
}

td {  
    background: #FAFAFA;
    text-align: center;
}

</style>
