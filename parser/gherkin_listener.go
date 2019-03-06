// Generated from Gherkin.g4 by ANTLR 4.7.

package parser // Gherkin

import "github.com/antlr/antlr4/runtime/Go/antlr"

// GherkinListener is a complete listener for a parse tree produced by GherkinParser.
type GherkinListener interface {
	antlr.ParseTreeListener

	// EnterTagName is called when entering the tagName production.
	EnterTagName(c *TagNameContext)

	// EnterTag is called when entering the tag production.
	EnterTag(c *TagContext)

	// EnterLineContent is called when entering the lineContent production.
	EnterLineContent(c *LineContentContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterTrailingComment is called when entering the trailingComment production.
	EnterTrailingComment(c *TrailingCommentContext)

	// EnterDescriptionLine is called when entering the descriptionLine production.
	EnterDescriptionLine(c *DescriptionLineContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterFeature is called when entering the feature production.
	EnterFeature(c *FeatureContext)

	// EnterBackground is called when entering the background production.
	EnterBackground(c *BackgroundContext)

	// EnterAbstractScenario is called when entering the abstractScenario production.
	EnterAbstractScenario(c *AbstractScenarioContext)

	// EnterScenario is called when entering the scenario production.
	EnterScenario(c *ScenarioContext)

	// EnterOutline is called when entering the outline production.
	EnterOutline(c *OutlineContext)

	// EnterExamples is called when entering the examples production.
	EnterExamples(c *ExamplesContext)

	// EnterStep is called when entering the step production.
	EnterStep(c *StepContext)

	// EnterRow is called when entering the row production.
	EnterRow(c *RowContext)

	// EnterCell is called when entering the cell production.
	EnterCell(c *CellContext)

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterDocumentIndent is called when entering the documentIndent production.
	EnterDocumentIndent(c *DocumentIndentContext)

	// EnterDocumentContent is called when entering the documentContent production.
	EnterDocumentContent(c *DocumentContentContext)

	// ExitTagName is called when exiting the tagName production.
	ExitTagName(c *TagNameContext)

	// ExitTag is called when exiting the tag production.
	ExitTag(c *TagContext)

	// ExitLineContent is called when exiting the lineContent production.
	ExitLineContent(c *LineContentContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitTrailingComment is called when exiting the trailingComment production.
	ExitTrailingComment(c *TrailingCommentContext)

	// ExitDescriptionLine is called when exiting the descriptionLine production.
	ExitDescriptionLine(c *DescriptionLineContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitFeature is called when exiting the feature production.
	ExitFeature(c *FeatureContext)

	// ExitBackground is called when exiting the background production.
	ExitBackground(c *BackgroundContext)

	// ExitAbstractScenario is called when exiting the abstractScenario production.
	ExitAbstractScenario(c *AbstractScenarioContext)

	// ExitScenario is called when exiting the scenario production.
	ExitScenario(c *ScenarioContext)

	// ExitOutline is called when exiting the outline production.
	ExitOutline(c *OutlineContext)

	// ExitExamples is called when exiting the examples production.
	ExitExamples(c *ExamplesContext)

	// ExitStep is called when exiting the step production.
	ExitStep(c *StepContext)

	// ExitRow is called when exiting the row production.
	ExitRow(c *RowContext)

	// ExitCell is called when exiting the cell production.
	ExitCell(c *CellContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitDocumentIndent is called when exiting the documentIndent production.
	ExitDocumentIndent(c *DocumentIndentContext)

	// ExitDocumentContent is called when exiting the documentContent production.
	ExitDocumentContent(c *DocumentContentContext)
}
