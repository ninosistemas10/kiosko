package category

import (
	"log"

	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ninosistemas10/kiosko/config"
	"github.com/ninosistemas10/kiosko/domain/category"
	"github.com/ninosistemas10/kiosko/infrastructure/handler/response"
	"github.com/ninosistemas10/kiosko/model"
)

type handler struct {
	useCase  category.UseCase
	response response.API
}

func newHandler(useCase category.UseCase) handler {
	return handler{useCase: useCase}
}

func (h handler) Create(c echo.Context) error {
	m := model.Category{}
	if err := c.Bind(&m); err != nil {
		return h.response.BindFailed(err)
	}

	if err := h.useCase.Create(&m); err != nil {
		return h.response.Error(c, "useCase.Create()", err)
	}

	return c.JSON(h.response.Created(m))
}

func (h handler) Update(c echo.Context) error {
	m := model.Category{}
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

	file, err := c.FormFile("image")
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
		Folder:   "categorias", // Carpeta en Cloudinary
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
		"images":  imageURL,
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
		h.response.Error(c, "uuid.Parse()", err)
	}

	categoryData, err := h.useCase.GetByID(ID)
	if err != nil {
		return h.response.Error(c, "useCase.GetBYID", err)
	}

	return c.JSON(h.response.OK(categoryData))
}

func (h handler) GetAll(c echo.Context) error {
	categorys, err := h.useCase.GetAll()
	if err != nil {
		return h.response.Error(c, "useCase.GetAll", err)
	}

	return c.JSON(h.response.OK(categorys))
}
