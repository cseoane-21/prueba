package prueba

async function controller2(req, res) {
const conn = await getConnection()
// hola
conn.query("SELECT * FROM Post WHERE id = ?", req.params.id);
}