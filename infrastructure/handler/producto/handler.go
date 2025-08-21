package producto

import (
	"fmt"
	"log"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/kiosko/config"
	"github.com/ninosistemas10/kiosko/domain/producto"
	"github.com/ninosistemas10/kiosko/model"

	"github.com/ninosistemas10/kiosko/infrastructure/handler/response"
)

type handler struct {
	useCase  producto.UseCase
	response response.API
}

func newHandler(useCase producto.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create(c echo.Context) error {
	m := model.Producto{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

func (h handler) Update(c echo.Context) error {
	m := model.Producto{}

	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	m.ID = ID

	if err := h.useCase.Update(&m); err != nil {
		return h.response.Error(c, "h.useCase.Update()", err)
	}

	return c.JSON(h.response.Updated(m))
}

func (h handler) UpdateImage(c echo.Context) error {
	log.Println("🚀 UpdateImage endpoint called")
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	file, err := c.FormFile("imagen")
	if err != nil {
		return h.response.Error(c, "No image file provided", err)
	}

	src, err := file.Open()
	if err != nil {
		return h.response.Error(c, "Unable to open image file", err)
	}
	defer src.Close()

	// Configuración de Cloudinary
	cld := config.SetupCloudinary()

	// Crear un nombre único para la imagen
	filename := uuid.New().String() + "_" + file.Filename

	// Subir a Cloudinary
	uploadResult, err := cld.Upload.Upload(c.Request().Context(), src, uploader.UploadParams{
		Folder:   "products", // Carpeta en Cloudinary
		PublicID: filename,
	})
	if err != nil {
		return h.response.Error(c, "Error uploading image to Cloudinary", err)
	}

	// Obtener la URL segura
	imageURL := uploadResult.SecureURL

	// Actualizar la URL en la base de datos
	err = h.useCase.UpdateImage(ID, imageURL) // <- Aquí debe actualizarse la URL en la DB
	if err != nil {
		return h.response.Error(c, "Error updating image URL in database", err)
	}

	return c.JSON(h.response.OK(map[string]string{
		"message": "Image updated successfully",
		"imagen":  imageURL,
	}))
}

func (h handler) Delete(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.BindFailed(err)
	}

	err = h.useCase.Delete(ID)
	if err != nil {
		return h.response.Error(c, "useCase.Delete()", err)
	}

	return c.JSON(h.response.Deleted(nil))
}

func (h handler) GetByID(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return h.response.Error(c, "uuid.Parse()", err)
	}

	productoData, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.response.Error(c, "useCase.GetWhere()", err)
	}

	return c.JSON(h.response.OK(productoData))
}

func (h handler) GetByCategoryID(c echo.Context) error {
	idCategoria, err := uuid.Parse(c.Param("idcategoria"))
	fmt.Println("Valor importante de idcategoria:", idCategoria)

	// Imprime el contexto
	fmt.Printf("Contexto: %+v\n", c)

	if err != nil {
		return h.response.Error(c, "uuid.Parse()", err)
	}

	productos, err := h.useCase.GetByCategoryID(idCategoria)
	if err != nil {
		return h.response.Error(c, "useCase.GetByCategoryID", err)
	}

	return c.JSON(h.response.OK(productos))
}

func (h handler) GetAll(c echo.Context) error {
	productos, err := h.useCase.GetAll()
	if err != nil {
		return h.response.Error(c, "useCase.GetAllWhere()", err)
	}

	return c.JSON(h.response.OK(productos))
}
