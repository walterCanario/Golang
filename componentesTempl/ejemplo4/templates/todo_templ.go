// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
// components/todo.templ

package templates

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func TodoApp() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"todo-container\"><h1>Todo List</h1><!-- Contador de tareas --><div id=\"task-counter\">Tareas completadas: <span id=\"completed-count\">0</span> / <span id=\"total-count\">0</span></div><!-- Formulario para agregar tareas --><form id=\"todo-form\" class=\"mb-4\"><input type=\"text\" id=\"new-task\" class=\"border p-2 mr-2\" placeholder=\"Nueva tarea\" required> <button type=\"submit\" class=\"bg-blue-500 text-white px-4 py-2 rounded\">Agregar</button></form><!-- Lista de tareas --><ul id=\"todo-list\" class=\"space-y-2\"></ul></div><script>\r\n        // Estado de la aplicación\r\n        let state = {\r\n            tasks: [],\r\n            completedTasks: 0\r\n        };\r\n\r\n        // Elementos del DOM\r\n        const form = document.getElementById('todo-form');\r\n        const input = document.getElementById('new-task');\r\n        const list = document.getElementById('todo-list');\r\n        const completedCount = document.getElementById('completed-count');\r\n        const totalCount = document.getElementById('total-count');\r\n\r\n        // Manejador para agregar tareas\r\n        form.addEventListener('submit', (e) => {\r\n            e.preventDefault();\r\n            \r\n            const taskText = input.value.trim();\r\n            if (taskText) {\r\n                addTask(taskText);\r\n                input.value = '';\r\n            }\r\n        });\r\n\r\n        // Función para agregar una tarea\r\n        function addTask(text) {\r\n            const task = {\r\n                id: Date.now(),\r\n                text,\r\n                completed: false\r\n            };\r\n\r\n            state.tasks.push(task);\r\n            \r\n            const li = document.createElement('li');\r\n            li.className = 'flex items-center space-x-2 p-2 border rounded';\r\n            li.innerHTML = `\r\n                <input \r\n                    type=\"checkbox\" \r\n                    class=\"form-checkbox\"\r\n                    onchange=\"toggleTask(${task.id})\"\r\n                />\r\n                <span class=\"flex-grow\">${text}</span>\r\n                <button \r\n                    onclick=\"deleteTask(${task.id})\"\r\n                    class=\"text-red-500 hover:text-red-700\"\r\n                >\r\n                    Eliminar\r\n                </button>\r\n            `;\r\n\r\n            list.appendChild(li);\r\n            updateCounters();\r\n        }\r\n\r\n        // Función para alternar el estado de una tarea\r\n        function toggleTask(id) {\r\n            const task = state.tasks.find(t => t.id === id);\r\n            if (task) {\r\n                task.completed = !task.completed;\r\n                state.completedTasks += task.completed ? 1 : -1;\r\n                updateCounters();\r\n            }\r\n        }\r\n\r\n        // Función para eliminar una tarea\r\n        function deleteTask(id) {\r\n            const index = state.tasks.findIndex(t => t.id === id);\r\n            if (index !== -1) {\r\n                const task = state.tasks[index];\r\n                if (task.completed) {\r\n                    state.completedTasks--;\r\n                }\r\n                state.tasks.splice(index, 1);\r\n                list.children[index].remove();\r\n                updateCounters();\r\n            }\r\n        }\r\n\r\n        // Función para actualizar los contadores\r\n        function updateCounters() {\r\n            completedCount.textContent = state.completedTasks;\r\n            totalCount.textContent = state.tasks.length;\r\n        }\r\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
